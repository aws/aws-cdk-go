package awsapigatewayv2


// Attributes for importing a WebSocketApi into the CDK.
//
// Example:
//   webSocketApi := apigwv2.WebSocketApi_FromWebSocketApiAttributes(this, jsii.String("mywsapi"), &WebSocketApiAttributes{
//   	WebSocketId: jsii.String("api-1234"),
//   })
//
type WebSocketApiAttributes struct {
	// The identifier of the WebSocketApi.
	WebSocketId *string `field:"required" json:"webSocketId" yaml:"webSocketId"`
	// The endpoint URL of the WebSocketApi.
	// Default: - throw san error if apiEndpoint is accessed.
	//
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

