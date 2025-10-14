package awsapigatewayv2


// Props for WebSocket API.
//
// Example:
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &WebSocketApiProps{
//   	ApiKeySelectionExpression: apigwv2.WebSocketApiKeySelectionExpression_HEADER_X_API_KEY(),
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
	// Avoid validating models when creating a deployment.
	// Default: false.
	//
	DisableSchemaValidation *bool `field:"optional" json:"disableSchemaValidation" yaml:"disableSchemaValidation"`
	// Options to configure a '$disconnect' route.
	// Default: - no '$disconnect' route configured.
	//
	DisconnectRouteOptions *WebSocketRouteOptions `field:"optional" json:"disconnectRouteOptions" yaml:"disconnectRouteOptions"`
	// The IP address types that can invoke the API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-ip-address-type.html
	//
	// Default: undefined - AWS default is IPV4.
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The route selection expression for the API.
	// Default: '$request.body.action'
	//
	RouteSelectionExpression *string `field:"optional" json:"routeSelectionExpression" yaml:"routeSelectionExpression"`
}

