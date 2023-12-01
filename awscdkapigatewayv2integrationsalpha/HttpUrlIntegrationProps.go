package awscdkapigatewayv2integrationsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
)

// Properties to initialize a new `HttpProxyIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//   import apigatewayv2_integrations_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var parameterMapping parameterMapping
//
//   httpUrlIntegrationProps := &HttpUrlIntegrationProps{
//   	Method: apigatewayv2_alpha.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   }
//
// Deprecated.
type HttpUrlIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Default: HttpMethod.ANY
	//
	// Deprecated.
	Method awscdkapigatewayv2alpha.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	// Deprecated.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
}

