package interfaces

import (
	"net/http"
)

// HTTPConfiguration encapsulates top-level HTTP configuration that applies to all SDK components.
//
// See ldcomponents.HTTPConfigurationBuilder for more details on these properties.
type HTTPConfiguration struct {
	// DefaultHeaders contains the basic headers that should be added to all HTTP requests from SDK
	// components to LaunchDarkly services, based on the current SDK configuration. This map is never
	// modified once created.
	DefaultHeaders http.Header

	// CreateHTTPClient is a function that returns a new HTTP client instance based on the SDK configuration.
	//
	// The SDK will ensure that this field is non-nil before passing it to any component.
	CreateHTTPClient func() *http.Client
}

// HTTPConfigurationFactory is an interface for a factory that creates an HTTPConfiguration.
type HTTPConfigurationFactory interface {
	// CreateHTTPConfiguration is called internally by the SDK to obtain the configuration.
	//
	// This happens only when MakeClient or MakeCustomClient is called. If the factory returns
	// an error, creation of the LDClient fails.
	CreateHTTPConfiguration(basicConfig BasicConfiguration) (HTTPConfiguration, error)
}
