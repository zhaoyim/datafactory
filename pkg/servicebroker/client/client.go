package client

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	"io/ioutil"
	"net/http"
)

type ServiceList struct {
	Services []backingserviceapi.BackingServiceSpec `json:"services"`
}

type Interface interface {
	Catalog(Url string, credential ...string) (ServiceList, error)
}

func NewServiceBrokerClient() Interface {
	return &httpClient{
		Get:  httpGet,
		Post: httpPostJson,
	}
}

type httpClient struct {
	Get  func(getUrl string, credential ...string) ([]byte, error)
	Post func(getUrl string, body []byte, credential ...string) ([]byte, error)
}

func (c *httpClient) Catalog(Url string, credential ...string) (ServiceList, error) {
	services := new(ServiceList)
	b, err := c.Get(Url+"/v2/catalog", credential...)
	if err != nil {
		fmt.Printf("httpclient catalog err %s", err.Error())
		return *services, err
	}

	if err := json.Unmarshal(b, services); err != nil {
		return *services, err
	}

	return *services, nil
}

//todo 支持多种自定义认证方式
func httpGet(getUrl string, credential ...string) ([]byte, error) {
	var resp *http.Response
	var err error
	if len(credential) == 2 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest("GET", getUrl, nil)
		if err != nil {
			return nil, fmt.Errorf("[servicebroker http client] err %s, %s\n", getUrl, err)
		}

		basic := fmt.Sprintf("Basic %s", string(base64Encode([]byte(fmt.Sprintf("%s:%s", credential[0], credential[1])))))
		req.Header.Set(Authorization, basic)

		resp, err = client.Do(req)
		if err != nil {
			fmt.Errorf("http get err:%s", err.Error())
			return nil, err
		}
		defer resp.Body.Close()
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

	glog.Infof("GET %s returns http code %v", getUrl, resp.StatusCode)
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
	if len(credential) == 2 {
		basic := fmt.Sprintf("Basic %s", string(base64Encode([]byte(fmt.Sprintf("%s:%s", credential[0], credential[1])))))
		req.Header.Set(Authorization, basic)
	}
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

const Authorization = "Authorization"

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}
