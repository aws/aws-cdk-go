package awsapigatewayv2


// The attributes used to import existing WebSocketStage.
//
// Example:
//   var webSocketApi IWebSocketApi
//
//
//   importedStage := apigwv2.WebSocketStage_FromWebSocketStageAttributes(this, jsii.String("imported-stage"), &WebSocketStageAttributes{
//   	StageName: jsii.String("myStage"),
//   	Api: webSocketApi,
//   })
//
//   apiKey := apigwv2.NewApiKey(this, jsii.String("MyApiKey"))
//
//   usagePlan := apigwv2.NewUsagePlan(this, jsii.String("MyUsagePlan"), &UsagePlanProps{
//   	ApiStages: []UsagePlanPerApiStage{
//   		&UsagePlanPerApiStage{
//   			Api: webSocketApi,
//   			Stage: importedStage,
//   		},
//   	},
//   })
//
//   usagePlan.addApiKey(apiKey)
//
type WebSocketStageAttributes struct {
	// The name of the stage.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// The API to which this stage is associated.
	Api IWebSocketApi `field:"required" json:"api" yaml:"api"`
}

