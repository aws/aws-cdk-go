package awsapigatewayv2


// Props for WebSocket API.
//
// Example:
//   import "github.com/aws-samples/dummy/awscdklib/awsapigatewayv2integrations"
//
//   var connectHandler function
//   var disconnectHandler function
//   var defaultHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &WebSocketApiProps{
//   	ConnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: awscdklibawsapigatewayv2integrations.NewWebSocketLambdaIntegration(jsii.String("ConnectIntegration"), connectHandler),
//   	},
//   	DisconnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: *awscdklibawsapigatewayv2integrations.NewWebSocketLambdaIntegration(jsii.String("DisconnectIntegration"), disconnectHandler),
//   	},
//   	DefaultRouteOptions: &WebSocketRouteOptions{
//   		Integration: *awscdklibawsapigatewayv2integrations.NewWebSocketLambdaIntegration(jsii.String("DefaultIntegration"), defaultHandler),
//   	},
//   })
//
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//
type WebSocketApiProps struct {
	// An API key selection expression.
	//
	// Providing this option will require an API Key be provided to access the API.
	// Default: - Key is not required to access these APIs.
	//
	ApiKeySelectionExpression WebSocketApiKeySelectionExpression `field:"optional" json:"apiKeySelectionExpression" yaml:"apiKeySelectionExpression"`
	// Name for the WebSocket API resource.
	// Default: - id of the WebSocketApi construct.
	//
	ApiName *string `field:"optional" json:"apiName" yaml:"apiName"`
	// Options to configure a '$connect' route.
	// Default: - no '$connect' route configured.
	//
	ConnectRouteOptions *WebSocketRouteOptions `field:"optional" json:"connectRouteOptions" yaml:"connectRouteOptions"`
	// Options to configure a '$default' route.
	// Default: - no '$default' route configured.
	//
	DefaultRouteOptions *WebSocketRouteOptions `field:"optional" json:"defaultRouteOptions" yaml:"defaultRouteOptions"`
	// The description of the API.
	// Default: - none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Options to configure a '$disconnect' route.
	// Default: - no '$disconnect' route configured.
	//
	DisconnectRouteOptions *WebSocketRouteOptions `field:"optional" json:"disconnectRouteOptions" yaml:"disconnectRouteOptions"`
	// The route selection expression for the API.
	// Default: '$request.body.action'
	//
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
}

