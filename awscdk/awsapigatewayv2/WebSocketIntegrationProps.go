package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var webSocketApi webSocketApi
//
//   webSocketIntegrationProps := &WebSocketIntegrationProps{
//   	IntegrationType: awscdk.Aws_apigatewayv2.WebSocketIntegrationType_AWS_PROXY,
//   	IntegrationUri: jsii.String("integrationUri"),
//   	WebSocketApi: webSocketApi,
//
//   	// the properties below are optional
//   	CredentialsRole: role,
//   	IntegrationMethod: jsii.String("integrationMethod"),
//   	PassthroughBehavior: awscdk.*Aws_apigatewayv2.PassthroughBehavior_WHEN_NO_MATCH,
//   	RequestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	RequestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	TemplateSelectionExpression: jsii.String("templateSelectionExpression"),
//   }
//
type WebSocketIntegrationProps struct {
	// Integration type.
	IntegrationType WebSocketIntegrationType `field:"required" json:"integrationType" yaml:"integrationType"`
	// Integration URI.
	IntegrationUri *string `field:"required" json:"integrationUri" yaml:"integrationUri"`
	// The WebSocket API to which this integration should be bound.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
	// Specifies the IAM role required for the integration.
	// Default: - No IAM role required.
	//
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// Specifies the integration's HTTP method type.
	// Default: - No HTTP method required.
	//
	IntegrationMethod *string `field:"optional" json:"integrationMethod" yaml:"integrationMethod"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	// Default: - No passthrough behavior required.
	//
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	// Default: - No request parameters required.
	//
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// ```
	//   { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// Default: - No request templates required.
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The template selection expression for the integration.
	// Default: - No template selection expression required.
	//
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

