package awscdkapigatewayv2alpha


// The attributes used to import existing WebSocketStage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var webSocketApi webSocketApi
//
//   webSocketStageAttributes := &WebSocketStageAttributes{
//   	Api: webSocketApi,
//   	StageName: jsii.String("stageName"),
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

