package swiftype

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	VERSION               = "1.0.0"
	USER_AGENT            = "Swiftype-Go/" + VERSION
	DEFAULT_API_HOST      = "api.swiftype.com"
	DEFAULT_API_BASE_PATH = "/api/v1/"
)

type JsonObject map[string]interface{}

type JsonArray []JsonObject

type SwiftypeSearchResults struct {
	Records JsonObject `json:"records"`
	Info    JsonObject `json:"info"`
	Errors  JsonObject `json:"errors"`
}

type Client struct {
	username string
	password string
	api_key  string
	host     string
	httpc    *http.Client
}

func NewClientWithUsernamePassword(username string, password string, host string) *Client {
	return &Client{username: username, password: password, host: host, httpc: &http.Client{}}
}

func NewClientWithApiKey(api_key string, host string) *Client {
	return &Client{api_key: api_key, host: host, httpc: &http.Client{}}
}

func (c *Client) decode(resp *http.Response) []byte {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		panic(err)
	}

	return b
}

func (c *Client) newRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)

	if err != nil {
		return
	}

	if len(c.username) != 0 && len(c.password) != 0 {
		req.SetBasicAuth(c.username, c.password)
	} else {
		req.Header.Add("Authorization", fmt.Sprintf("auth_token %s", c.api_key))
	}

	return
}

// func (c *Client) executeRequest(req *http.Request) (res JsonObject, err error) {

// 	resp, err := c.httpc.Do(req)

// 	if err != nil {
// 		return
// 	}

// 	return

// }

func (c *Client) get(path string, params url.Values) []byte {

	params.Add("auth_token", c.api_key)

	url := fmt.Sprintf("https://%s%s.json?%s", c.host, path, params.Encode())

	fmt.Println(url, "BBBB")

	resp, err := c.httpc.Get(url)

	if err != nil {
		panic(err)
	}

	return c.decode(resp)
}

func (c *Client) delete(path string, params url.Values) JsonArray {
	return nil
}

func (c *Client) put(path string, params url.Values) JsonArray {
	return nil
}

func (c *Client) post(path string, params url.Values) JsonArray {
	return nil
}

func (c *Client) Engines() []byte {

	results := c.get(DEFAULT_API_BASE_PATH+"engines", url.Values{})

	return results
}

func (c *Client) Engine(engine string) interface{} {

	params := url.Values{}
	params.Set("name", engine)

	results := c.get(DEFAULT_API_BASE_PATH+"engines", params)

	return results
}

func (c *Client) Search(engine string, query string) []byte {

	params := url.Values{}
	params.Set("q", query)

	path := fmt.Sprintf("%sengines/%s/search", DEFAULT_API_BASE_PATH, engine)

	results := c.get(path, params)

	return results
}
