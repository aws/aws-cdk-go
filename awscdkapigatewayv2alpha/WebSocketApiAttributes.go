package awscdkapigatewayv2alpha


// Attributes for importing a WebSocketApi into the CDK.
//
// Example:
//   webSocketApi := apigwv2.WebSocketApi_FromWebSocketApiAttributes(this, jsii.String("mywsapi"), &WebSocketApiAttributes{
//   	WebSocketId: jsii.String("api-1234"),
//   })
//
// Experimental.
type WebSocketApiAttributes struct {
	// The identifier of the WebSocketApi.
	// Experimental.
	WebSocketId *string `field:"required" json:"webSocketId" yaml:"webSocketId"`
	// The endpoint URL of the WebSocketApi.
	// Experimental.
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

