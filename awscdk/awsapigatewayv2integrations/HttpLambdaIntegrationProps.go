package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

// Lambda Proxy integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpLambdaIntegrationProps := &HttpLambdaIntegrationProps{
//   	ParameterMapping: parameterMapping,
//   	PayloadFormatVersion: payloadFormatVersion,
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type HttpLambdaIntegrationProps struct {
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Default: undefined requests are sent to the backend unmodified.
	//
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Default: PayloadFormatVersion.VERSION_2_0
	//
	PayloadFormatVersion awsapigatewayv2.PayloadFormatVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

