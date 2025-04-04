package awsapigatewayv2


// Container for defining throttling parameters to API stages.
//
// Example:
//   var api httpApi
//
//
//   apigwv2.NewHttpStage(this, jsii.String("Stage"), &HttpStageProps{
//   	HttpApi: api,
//   	Throttle: &ThrottleSettings{
//   		RateLimit: jsii.Number(1000),
//   		BurstLimit: jsii.Number(1000),
//   	},
//   	DetailedMetricsEnabled: jsii.Boolean(true),
//   })
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

