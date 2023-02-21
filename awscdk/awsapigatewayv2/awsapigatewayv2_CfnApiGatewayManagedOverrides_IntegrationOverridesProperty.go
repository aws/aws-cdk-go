package awsapigatewayv2


// The `IntegrationOverrides` property overrides the integration settings for an API Gateway-managed integration.
//
// If you remove this property, API Gateway restores the default values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationOverridesProperty := &integrationOverridesProperty{
//   	description: jsii.String("description"),
//   	integrationMethod: jsii.String("integrationMethod"),
//   	payloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	timeoutInMillis: jsii.Number(123),
//   }
//
type CfnApiGatewayManagedOverrides_IntegrationOverridesProperty struct {
	// The description of the integration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the integration's HTTP method type.
	IntegrationMethod *string `field:"optional" json:"integrationMethod" yaml:"integrationMethod"`
	// Specifies the format of the payload sent to an integration.
	//
	// Required for HTTP APIs. For HTTP APIs, supported values for Lambda proxy integrations are `1.0` and `2.0` . For all other integrations, `1.0` is the only supported value. To learn more, see [Working with AWS Lambda proxy integrations for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-lambda.html) .
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	//
	// The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
}

