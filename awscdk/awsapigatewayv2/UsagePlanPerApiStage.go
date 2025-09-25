package awsapigatewayv2


// Represents the API stages that a usage plan applies to.
//
// Example:
//   api := apigwv2.NewWebSocketApi(this, jsii.String("my-api"))
//   stage := apigwv2.NewWebSocketStage(this, jsii.String("my-stage"), &WebSocketStageProps{
//   	WebSocketApi: api,
//   	StageName: jsii.String("dev"),
//   })
//
//   usagePlan := apigwv2.NewUsagePlan(this, jsii.String("my-usage-plan"), &UsagePlanProps{
//   	UsagePlanName: jsii.String("Basic"),
//   })
//
//   usagePlan.AddApiStage(&UsagePlanPerApiStage{
//   	Api: api,
//   	Stage: stage,
//   })
//
type UsagePlanPerApiStage struct {
	// The WebSocket API to associate with the usage plan.
	// Default: none.
	//
	Api IWebSocketApi `field:"optional" json:"api" yaml:"api"`
	// [disable-awslint:ref-via-interface].
	// Default: none.
	//
	Stage IWebSocketStage `field:"optional" json:"stage" yaml:"stage"`
}

