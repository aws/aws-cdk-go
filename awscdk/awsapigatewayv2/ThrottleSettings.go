package awsapigatewayv2


// Container for defining throttling parameters to API stages.
//
// Example:
//   apiKey := apigwv2.NewApiKey(this, jsii.String("ApiKey"))
//
//   usagePlan := apigwv2.NewUsagePlan(this, jsii.String("UsagePlan"), &UsagePlanProps{
//   	UsagePlanName: jsii.String("WebSocketUsagePlan"),
//   	Throttle: &ThrottleSettings{
//   		RateLimit: jsii.Number(10),
//   		BurstLimit: jsii.Number(2),
//   	},
//   })
//
//   usagePlan.addApiKey(apiKey)
//
type ThrottleSettings struct {
	// The maximum API request rate limit over a time ranging from one to a few seconds.
	// Default: none.
	//
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// The API request steady-state rate limit (average requests per second over an extended period of time).
	// Default: none.
	//
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

