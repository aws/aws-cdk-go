// Integrations for AWS APIGateway V2
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
//   httpUrlIntegrationProps := &httpUrlIntegrationProps{
//   	method: apigatewayv2_alpha.httpMethod_ANY,
//   	parameterMapping: parameterMapping,
//   }
//
// Experimental.
type HttpUrlIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awscdkapigatewayv2alpha.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awscdkapigatewayv2alpha.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
}

