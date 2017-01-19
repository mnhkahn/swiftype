package swiftype

import (
	"encoding/json"
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

func (c *Client) decode(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
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

func (c *Client) get(path string, params url.Values) ([]byte, error) {
	params.Add("auth_token", c.api_key)

	url := fmt.Sprintf("https://%s%s.json?%s", c.host, path, params.Encode())

	resp, err := c.httpc.Get(url)

	if err != nil {
		return nil, err
	}

	return c.decode(resp)
}

// func (c *Client) delete(path string, params url.Values) JsonArray {
// 	return nil
// }

// func (c *Client) put(path string, params url.Values) JsonArray {
// 	return nil
// }

// func (c *Client) post(path string, params url.Values) JsonArray {
// 	return nil
// }

func (c *Client) Engines() ([]byte, error) {
	return c.get(DEFAULT_API_BASE_PATH+"engines", url.Values{})
}

func (c *Client) Engine(engine string) ([]byte, error) {
	params := url.Values{}
	params.Set("name", engine)

	return c.get(DEFAULT_API_BASE_PATH+"engines", params)
}

func (c *Client) Search(engine string, query string) (*SwiftypeResult, error) {
	params := url.Values{}
	params.Set("q", query)

	path := fmt.Sprintf("%sengines/%s/search", DEFAULT_API_BASE_PATH, engine)

	data, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	res := new(SwiftypeResult)
	err = json.Unmarshal(data, res)
	return res, err
}
