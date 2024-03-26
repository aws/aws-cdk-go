package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for AWS type integration for a WebSocket Api.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiRole role
//   var table table
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.AddRoute(jsii.String("$connect"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketAwsIntegration(jsii.String("DynamodbPutItem"), &WebSocketAwsIntegrationProps{
//   		IntegrationUri: fmt.Sprintf("arn:aws:apigateway:%v:dynamodb:action/PutItem", this.Region),
//   		IntegrationMethod: apigwv2.HttpMethod_POST,
//   		CredentialsRole: apiRole,
//   		RequestTemplates: map[string]*string{
//   			"application/json": JSON.stringify(map[string]interface{}{
//   				"TableName": table.tableName,
//   				"Item": map[string]map[string]*string{
//   					"id": map[string]*string{
//   						"S": jsii.String("$context.requestId"),
//   					},
//   				},
//   			}),
//   		},
//   	}),
//   })
//
type WebSocketAwsIntegrationProps struct {
	// Specifies the integration's HTTP method type.
	IntegrationMethod *string `field:"required" json:"integrationMethod" yaml:"integrationMethod"`
	// Integration URI.
	IntegrationUri *string `field:"required" json:"integrationUri" yaml:"integrationUri"`
	// Specifies how to handle response payload content type conversions.
	// Default: - The response payload will be passed through from the integration response to
	// the route response or method response without modification.
	//
	ContentHandling awsapigatewayv2.ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Specifies the credentials role required for the integration.
	// Default: - No credential role provided.
	//
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	// Default: - No passthrough behavior required.
	//
	PassthroughBehavior awsapigatewayv2.PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	// Default: - No request parameter provided to the integration.
	//
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// ```
	//   { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// Default: - No request template provided to the integration.
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The template selection expression for the integration.
	// Default: - No template selection expression provided.
	//
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

