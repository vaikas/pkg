/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	"github.com/knative/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type DuckV1alpha1Interface interface {
	RESTClient() rest.Interface
	ChannelsGetter
	GenerationalsGetter
	KResourcesGetter
	LegacyTargetsGetter
	SinksGetter
	SubscriptionsGetter
	TargetsGetter
}

// DuckV1alpha1Client is used to interact with features provided by the duck.knative.dev group.
type DuckV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DuckV1alpha1Client) Channels(namespace string) ChannelInterface {
	return newChannels(c, namespace)
}

func (c *DuckV1alpha1Client) Generationals(namespace string) GenerationalInterface {
	return newGenerationals(c, namespace)
}

func (c *DuckV1alpha1Client) KResources(namespace string) KResourceInterface {
	return newKResources(c, namespace)
}

func (c *DuckV1alpha1Client) LegacyTargets(namespace string) LegacyTargetInterface {
	return newLegacyTargets(c, namespace)
}

func (c *DuckV1alpha1Client) Sinks(namespace string) SinkInterface {
	return newSinks(c, namespace)
}

func (c *DuckV1alpha1Client) Subscriptions(namespace string) SubscriptionInterface {
	return newSubscriptions(c, namespace)
}

func (c *DuckV1alpha1Client) Targets(namespace string) TargetInterface {
	return newTargets(c, namespace)
}

// NewForConfig creates a new DuckV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*DuckV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DuckV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new DuckV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DuckV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DuckV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *DuckV1alpha1Client {
	return &DuckV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *DuckV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
