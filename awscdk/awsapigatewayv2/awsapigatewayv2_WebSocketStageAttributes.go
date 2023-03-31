package awsapigatewayv2


// The attributes used to import existing WebSocketStage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var webSocketApi webSocketApi
//
//   webSocketStageAttributes := &webSocketStageAttributes{
//   	api: webSocketApi,
//   	stageName: jsii.String("stageName"),
//   }
//
// Experimental.
type WebSocketStageAttributes struct {
	// The name of the stage.
	// Experimental.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The API to which this stage is associated.
	// Experimental.
	Api IWebSocketApi `field:"required" json:"api" yaml:"api"`
}

