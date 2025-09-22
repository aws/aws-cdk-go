package awsapigatewayv2


// Properties for defining an API Gateway Usage Plan for WebSocket APIs.
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
type UsagePlanProps struct {
	// API Stages to be associated with the usage plan.
	// Default: none.
	//
	ApiStages *[]*UsagePlanPerApiStage `field:"optional" json:"apiStages" yaml:"apiStages"`
	// Represents usage plan purpose.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Number of requests clients can make in a given time period.
	// Default: none.
	//
	Quota *QuotaSettings `field:"optional" json:"quota" yaml:"quota"`
	// Overall throttle settings for the API.
	// Default: none.
	//
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
	// Name for this usage plan.
	// Default: none.
	//
	UsagePlanName *string `field:"optional" json:"usagePlanName" yaml:"usagePlanName"`
}

