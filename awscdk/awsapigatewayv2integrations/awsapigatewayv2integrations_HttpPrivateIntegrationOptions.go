package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
)

// Base options for private integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//   var vpcLink vpcLink
//
//   httpPrivateIntegrationOptions := &HttpPrivateIntegrationOptions{
//   	Method: awscdk.Aws_apigatewayv2.HttpMethod_ANY,
//   	ParameterMapping: parameterMapping,
//   	SecureServerName: jsii.String("secureServerName"),
//   	VpcLink: vpcLink,
//   }
//
// Experimental.
type HttpPrivateIntegrationOptions struct {
	// The HTTP method that must be used to invoke the underlying HTTP proxy.
	// Experimental.
	Method awsapigatewayv2.HttpMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Specifies the server name to verified by HTTPS when calling the backend integration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-integration-tlsconfig.html
	//
	// Experimental.
	SecureServerName *string `field:"optional" json:"secureServerName" yaml:"secureServerName"`
	// The vpc link to be used for the private integration.
	// Experimental.
	VpcLink awsapigatewayv2.IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

