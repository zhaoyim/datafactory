package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	"io/ioutil"
	"net/http"
	"strings"
)

type ServiceList struct {
	Services []backingserviceapi.BackingServiceSpec `json:"services"`
}

type CreateServiceInstanceResponse struct {
	DashboardUrl  string         `json:"dashboard_url"`
	LastOperation *LastOperation `json:"last_operation, omitempty"`
}

type LastOperation struct {
	State                    string `json:"state"`
	Description              string `json:"description"`
	AsyncPollIntervalSeconds int    `json:"async_poll_interval_seconds, omitempty"`
}

/*
type Interface interface {
	Catalog(Url string) (ServiceList, error)
	CreateInstance() (interface{}, error)
}

func NewServiceBrokerClient() Interface {
	return &sbClient{
		Get:  httpGet,
		Post: httpPostJson,
	}
}

type sbClient struct {
	ServiceBroker struct {
		name     string
		url      string
		user     string
		password string
	}
	Get  func(getUrl string, credential ...string) ([]byte, error)
	Post func(getUrl string, body []byte, credential ...string) ([]byte, error)
}

func commToServiceBroker(method, path string, jsonData []byte, header map[string]string) (resp *http.Response, err error) {
	//fmt.Println(method, path, string(jsonData))

	req, err := http.NewRequest(strings.ToUpper(method) , path, bytes.NewBuffer(jsonData))

	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	return http.DefaultClient.Do(req)
}

func basicAuthStr(username, password string) string {
	auth := username + ":" + password
	authstr := base64.StdEncoding.EncodeToString([]byte(auth))
	return "Basic " + authstr
}

func (c *sbClient) Catalog(Url string) (ServiceList, error) {
	services := new(ServiceList)
	b, err := c.Get("http://" + Url + "/v2/catalog")
	if err != nil {
		fmt.Printf("httpclient catalog err %s", err.Error())
		return *services, err
	}

	if err := json.Unmarshal(b, services); err != nil {
		return *services, err
	}

	return *services, nil
}

func httpGet(getUrl string, credential ...string) ([]byte, error) {
	var resp *http.Response
	var err error
	if len(credential) == 2 {
		req, err := http.NewRequest("GET", getUrl, nil)
		if err != nil {
			return nil, fmt.Errorf("[servicebroker http client] err %s, %s\n", getUrl, err)
		}
		req.Header.Set(credential[0], credential[1])
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			fmt.Errorf("http get err:%s", err.Error())
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("[servicebroker http client] status err %s, %d\n", getUrl, resp.StatusCode)
		}
	} else {
		resp, err = http.Get(getUrl)
		if err != nil {
			fmt.Errorf("servicebroker http client get err:%s", err.Error())
			return nil, err
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("[http get] status err %s, %d\n", getUrl, resp.StatusCode)
		}
	}

	return ioutil.ReadAll(resp.Body)
}

func httpPostJson(postUrl string, body []byte, credential ...string) ([]byte, error) {
	var resp *http.Response
	var err error
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("[http] err %s, %s\n", postUrl, err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(credential[0], credential[1])
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("[http] err %s, %s\n", postUrl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("[http] status err %s, %d\n", postUrl, resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[http] read err %s, %s\n", postUrl, err)
	}
	return b, nil
}
*/
