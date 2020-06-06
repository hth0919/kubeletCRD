package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KetiV1Interface interface {
	RESTClient() rest.Interface
	PodsGetter
}

type KetiV1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*KetiV1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: "keti.migration", Version: "v1"}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()
	var cs KetiV1Client
	var err error

	cs.restClient, err = rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}
func NewForConfigOrDie(c *rest.Config) *KetiV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1Client for the given RESTClient.
func New(c rest.Interface) *KetiV1Client {
	return &KetiV1Client{c}
}

func (c *KetiV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *KetiV1Client) Pods(namespace string) PodInterface {
	return newPods(c, namespace)
}