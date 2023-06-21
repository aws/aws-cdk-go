package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// You can integrate an API method with an HTTP endpoint using the HTTP proxy integration or the HTTP custom integration,.
//
// With the proxy integration, the setup is simple. You only need to set the
// HTTP method and the HTTP endpoint URI, according to the backend requirements,
// if you are not concerned with content encoding or caching.
//
// With the custom integration, the setup is more involved. In addition to the
// proxy integration setup steps, you need to specify how the incoming request
// data is mapped to the integration request and how the resulting integration
// response data is mapped to the method response.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &RequestAuthorizerProps{
//   	Handler: authFn,
//   	IdentitySources: []*string{
//   		apigateway.IdentitySource_Header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	Authorizer: auth,
//   })
//
type HttpIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for HttpIntegration
type jsiiProxy_HttpIntegration struct {
	jsiiProxy_Integration
}

func NewHttpIntegration(url *string, props *HttpIntegrationProps) HttpIntegration {
	_init_.Initialize()

	if err := validateNewHttpIntegrationParameters(url, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.HttpIntegration",
		[]interface{}{url, props},
		&j,
	)

	return &j
}

func NewHttpIntegration_Override(h HttpIntegration, url *string, props *HttpIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.HttpIntegration",
		[]interface{}{url, props},
		h,
	)
}

func (h *jsiiProxy_HttpIntegration) Bind(_method Method) *IntegrationConfig {
	if err := h.validateBindParameters(_method); err != nil {
		panic(err)
	}
	var returns *IntegrationConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

