package awsapigatewayv2integrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Props for AWS type integration for an HTTP Api.
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
	// Specifies the credentials role required for the integration.
	// Default: - No credential role provided.
	//
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
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
}

