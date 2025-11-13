package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
)

// Lambda Proxy integration properties.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var booksDefaultFn Function
//
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   getBooksIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("GetBooksIntegration"), booksDefaultFn, &HttpLambdaIntegrationProps{
//   	ScopePermissionToRoute: jsii.Boolean(false),
//   })
//   createBookIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("CreateBookIntegration"), booksDefaultFn, &HttpLambdaIntegrationProps{
//   	ScopePermissionToRoute: jsii.Boolean(false),
//   })
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_GET,
//   	},
//   	Integration: getBooksIntegration,
//   })
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: createBookIntegration,
//   })
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
	// Scope the permission for invoking the AWS Lambda down to the specific route associated with this integration.
	//
	// If this is set to `false`, the permission will allow invoking the AWS Lambda
	// from any route. This is useful for reducing the AWS Lambda policy size
	// for cases where the same AWS Lambda function is reused for many integrations.
	// Default: true.
	//
	ScopePermissionToRoute *bool `field:"optional" json:"scopePermissionToRoute" yaml:"scopePermissionToRoute"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

