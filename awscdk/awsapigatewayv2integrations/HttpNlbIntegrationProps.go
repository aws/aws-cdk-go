package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

// Properties to initialize `HttpNlbIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping ParameterMapping
//   var vpcLink VpcLink
//
//   httpNlbIntegrationProps := &HttpNlbIntegrationProps{
//   	Method: awscdk.Aws_apigatewayv2.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   	SecureServerName: jsii.String("secureServerName"),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	VpcLink: vpcLink,
//   }
//
type HttpNlbIntegrationProps struct {
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
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Default: undefined private integration traffic will use HTTP protocol.
	//
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The vpc link to be used for the private integration.
	// Default: - a new VpcLink is created.
	//
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

