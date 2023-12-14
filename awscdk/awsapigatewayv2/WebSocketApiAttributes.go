package awsapigatewayv2


// Attributes for importing a WebSocketApi into the CDK.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketApiAttributes := &WebSocketApiAttributes{
//   	WebSocketId: jsii.String("webSocketId"),
//
//   	// the properties below are optional
//   	ApiEndpoint: jsii.String("apiEndpoint"),
//   }
//
type WebSocketApiAttributes struct {
	// The identifier of the WebSocketApi.
	WebSocketId *string `field:"required" json:"webSocketId" yaml:"webSocketId"`
	// The endpoint URL of the WebSocketApi.
	// Default: - throw san error if apiEndpoint is accessed.
	//
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
}

