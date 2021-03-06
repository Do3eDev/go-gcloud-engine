package go_gcloud_engine

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"path"
	"strconv"
)

// CustomerIO wraps the customer.io API, see: http://customer.io/docs/api/rest.html
type CustomerIO struct {
	siteID  string
	apiKey  string
	Host    string
	SSL     bool
	Request *http.Request
}

// CustomerIOError is returned by any method that fails at the API level
type CustomerIOError struct {
	status int
	url    string
	body   []byte
}

func (e *CustomerIOError) Error() string {
	return fmt.Sprintf("%v: %v %v", e.status, e.url, string(e.body))
}

// NewCustomerIO creates a new CustomerIO object to perform requests on the supplied credentials
func NewCustomerIO(siteID, apiKey string, request *http.Request) *CustomerIO {
	return &CustomerIO{siteID, apiKey, "track.customer.io", true, request}
}

// Identify identifies a customer and sets their attributes
func (c *CustomerIO) Identify(env string, customerID string, attributes map[string]interface{}) error {
	j, err := json.Marshal(attributes)

	if err != nil {
		return err
	}

	status, responseBody, err := c.request(env, "PUT", c.customerURL(customerID), j)

	if err != nil {
		return err
	} else if status != 200 {
		return &CustomerIOError{status, c.customerURL(customerID), responseBody}
	}

	return nil
}

// Track sends a single event to Customer.io for the supplied user
func (c *CustomerIO) Track(env string, customerID string, eventName string, data map[string]interface{}) error {

	body := map[string]interface{}{"name": eventName, "data": data}
	j, err := json.Marshal(body)

	if err != nil {
		return err
	}

	status, responseBody, err := c.request(env, "POST", c.eventURL(customerID), j)

	if err != nil {
		return err
	} else if status != 200 {
		return &CustomerIOError{status, c.eventURL(customerID), responseBody}
	}

	return nil
}

// TrackAnonymous sends a single event to Customer.io for the anonymous user
func (c *CustomerIO) TrackAnonymous(env string, eventName string, data map[string]interface{}) error {
	body := map[string]interface{}{"name": eventName, "data": data}
	j, err := json.Marshal(body)

	if err != nil {
		return err
	}

	status, responseBody, err := c.request(env, "POST", c.anonURL(), j)

	if err != nil {
		return err
	} else if status != 200 {
		return &CustomerIOError{status, c.anonURL(), responseBody}
	}

	return nil
}

// Delete deletes a customer
func (c *CustomerIO) Delete(env string, customerID string) error {
	status, responseBody, err := c.request(env, "DELETE", c.customerURL(customerID), []byte{})

	if err != nil {
		return err
	} else if status != 200 {
		return &CustomerIOError{status, c.customerURL(customerID), responseBody}
	}

	return nil
}

// AddDevice adds a device for a customer
func (c *CustomerIO) AddDevice(env string, customerID string, deviceID string, platform string, data map[string]interface{}) error {
	if customerID == "" {
		return errors.New("customerID is a required field")
	}
	if deviceID == "" {
		return errors.New("deviceID is a required field")
	}
	if platform == "" {
		return errors.New("platform is a required field")
	}

	body := map[string]map[string]interface{}{"device": {"id": deviceID, "platform": platform}}
	for k, v := range data {
		body["device"][k] = v
	}
	j, err := json.Marshal(body)

	if err != nil {
		return err
	}

	status, responseBody, err := c.request(env, "PUT", c.deviceURL(customerID), j)

	if err != nil {
		return err
	} else if status != 200 {
		return &CustomerIOError{status, c.deviceURL(customerID), responseBody}
	}

	return nil
}

// DeleteDevice deletes a device for a customer
func (c *CustomerIO) DeleteDevice(env string, customerID string, deviceID string) error {
	status, responseBody, err := c.request(env, "DELETE", c.deleteDeviceURL(customerID, deviceID), []byte{})

	if err != nil {
		return err
	} else if status != 200 {
		return &CustomerIOError{status, c.deleteDeviceURL(customerID, deviceID), responseBody}
	}

	return nil
}

func (c *CustomerIO) auth() string {
	return base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", c.siteID, c.apiKey)))
}

func (c *CustomerIO) protocol() string {
	if !c.SSL {
		return "http://"
	}
	return "https://"
}

func (c *CustomerIO) customerURL(customerID string) string {
	return c.protocol() + path.Join(c.Host, "api/v1", "customers", customerID)
}

func (c *CustomerIO) eventURL(customerID string) string {
	return c.protocol() + path.Join(c.Host, "api/v1", "customers", customerID, "events")
}

func (c *CustomerIO) anonURL() string {
	return c.protocol() + path.Join(c.Host, "api/v1", "events")
}

func (c *CustomerIO) deviceURL(customerID string) string {
	return c.protocol() + path.Join(c.Host, "api/v1", "customers", customerID, "devices")
}

func (c *CustomerIO) deleteDeviceURL(customerID string, deviceID string) string {
	return c.protocol() + path.Join(c.Host, "api/v1", "customers", customerID, "devices", deviceID)
}

func (c *CustomerIO) request(env, method, url string, body []byte) (status int, responseBody []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		return 0, nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %v", c.auth()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(body)))

	// Declare resp with Response type
	var resp *http.Response

	// If not local env, using url fetch of Google app engine
	if env != "local" {
		ctx := appengine.NewContext(c.Request)
		client := urlfetch.Client(ctx)
		resp, err = client.Do(req)
	} else {
		client := http.DefaultClient
		resp, err = client.Do(req)
	}

	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	status = resp.StatusCode

	if resp.ContentLength >= 0 {
		responseBody = make([]byte, resp.ContentLength)
		resp.Body.Read(responseBody)
	}

	return status, responseBody, nil
}
