package awsapigateway


// Example:
//   var integration lambdaIntegration
//
//
//   api := apigateway.NewRestApi(this, jsii.String("hello-api"))
//
//   v1 := api.Root.AddResource(jsii.String("v1"))
//   echo := v1.AddResource(jsii.String("echo"))
//   echoMethod := echo.AddMethod(jsii.String("GET"), integration, &MethodOptions{
//   	ApiKeyRequired: jsii.Boolean(true),
//   })
//
//   plan := api.AddUsagePlan(jsii.String("UsagePlan"), &UsagePlanProps{
//   	Name: jsii.String("Easy"),
//   	Throttle: &ThrottleSettings{
//   		RateLimit: jsii.Number(10),
//   		BurstLimit: jsii.Number(2),
//   	},
//   })
//
//   key := api.AddApiKey(jsii.String("ApiKey"))
//   plan.addApiKey(key)
//
type UsagePlanProps struct {
	// API Stages to be associated with the usage plan.
	ApiStages *[]*UsagePlanPerApiStage `field:"optional" json:"apiStages" yaml:"apiStages"`
	// Represents usage plan purpose.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name for this usage plan.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Number of requests clients can make in a given time period.
	Quota *QuotaSettings `field:"optional" json:"quota" yaml:"quota"`
	// Overall throttle settings for the API.
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
}

