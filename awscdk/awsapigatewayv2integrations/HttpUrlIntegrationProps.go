package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

// Properties to initialize a new `HttpProxyIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping ParameterMapping
//
//   httpUrlIntegrationProps := &HttpUrlIntegrationProps{
//   	Method: awscdk.Aws_apigatewayv2.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type HttpUrlIntegrationProps struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Default: HttpMethod.ANY
	//
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

