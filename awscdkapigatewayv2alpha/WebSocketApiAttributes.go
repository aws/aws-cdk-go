package awscdkapigatewayv2alpha


// Attributes for importing a WebSocketApi into the CDK.
//
// Example:
//   webSocketApi := apigwv2.WebSocketApi_FromWebSocketApiAttributes(this, jsii.String("mywsapi"), &WebSocketApiAttributes{
//   	WebSocketId: jsii.String("api-1234"),
//   })
//
// Deprecated.
type WebSocketApiAttributes struct {
	// The identifier of the WebSocketApi.
	// Deprecated.
	WebSocketId *string `field:"required" json:"webSocketId" yaml:"webSocketId"`
	// The endpoint URL of the WebSocketApi.
	// Default: - throw san error if apiEndpoint is accessed.
	//
	// Deprecated.
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

