package awsapigatewayv2


// Supported HTTP methods.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiRole Role
//   var table Table
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
type HttpMethod string

const (
	// HTTP ANY.
	HttpMethod_ANY HttpMethod = "ANY"
	// HTTP DELETE.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// HTTP GET.
	HttpMethod_GET HttpMethod = "GET"
	// HTTP HEAD.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// HTTP OPTIONS.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// HTTP PATCH.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// HTTP POST.
	HttpMethod_POST HttpMethod = "POST"
	// HTTP PUT.
	HttpMethod_PUT HttpMethod = "PUT"
)

