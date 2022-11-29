package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
)

// Lambda Proxy integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterMapping parameterMapping
//   var payloadFormatVersion payloadFormatVersion
//
//   httpLambdaIntegrationProps := &httpLambdaIntegrationProps{
//   	parameterMapping: parameterMapping,
//   	payloadFormatVersion: payloadFormatVersion,
//   }
//
// Experimental.
type HttpLambdaIntegrationProps struct {
	// Specifies how to transform HTTP requests before sending them to the backend.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html
	//
	// Experimental.
	ParameterMapping awsapigatewayv2.ParameterMapping `field:"optional" json:"parameterMapping" yaml:"parameterMapping"`
	// Version of the payload sent to the lambda handler.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html
	//
	// Experimental.
	PayloadFormatVersion awsapigatewayv2.PayloadFormatVersion `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
}

