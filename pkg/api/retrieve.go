package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

func (c *Client) RetrieveJson() ([]byte, error) {
	data, err := c.Retrieve()
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) Retrieve() (map[string]interface{}, error) {
	res, err := c.httpClient.PostForm(fmt.Sprintf("%s/retrieve", c.Options.Host), url.Values{
		"data": {"{\"op\": \"showConfig\", \"path\": []}"},
		"key":  {c.Options.ApiKey},
	})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}
	data := m["data"]
	d := data.(map[string]interface{})

	// delete login configuration under "system" due to complexities with encrypted passwords
	for key := range d {
		if key == "system" {
			users := d[key]
			u := users.(map[string]interface{})
			delete(u, "login")
		}
	}

	return d, nil
}
