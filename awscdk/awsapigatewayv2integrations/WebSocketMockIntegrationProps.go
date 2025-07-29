package awsapigatewayv2integrations


// Props for Mock type integration for a WebSocket Api.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//
//   webSocketApi.AddRoute(jsii.String("sendMessage"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketMockIntegration(jsii.String("DefaultIntegration"), &WebSocketMockIntegrationProps{
//   		RequestTemplates: map[string]*string{
//   			"application/json": JSON.stringify(map[string]*f64{
//   				"statusCode": jsii.Number(200),
//   			}),
//   		},
//   		TemplateSelectionExpression: jsii.String("\\$default"),
//   	}),
//   	ReturnResponse: jsii.Boolean(true),
//   })
//
type WebSocketMockIntegrationProps struct {
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// ```
	//   { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html
	//
	// Default: - No request template provided to the integration.
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The template selection expression for the integration.
	// Default: - No template selection expression provided.
	//
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

