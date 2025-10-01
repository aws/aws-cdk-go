package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

// Props for Lambda type integration for a WebSocket Api.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketLambdaIntegrationProps := &WebSocketLambdaIntegrationProps{
//   	ContentHandling: awscdk.Aws_apigatewayv2.ContentHandling_CONVERT_TO_BINARY,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type WebSocketLambdaIntegrationProps struct {
	// Specifies how to handle response payload content type conversions.
	// Default: - The response payload will be passed through from the integration response to
	// the route response or method response without modification.
	//
	ContentHandling awsapigatewayv2.ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

