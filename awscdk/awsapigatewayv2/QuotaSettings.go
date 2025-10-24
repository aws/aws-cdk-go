package awsapigatewayv2


// Specifies the maximum number of requests that clients can make to API Gateway APIs.
//
// Example:
//   var api WebSocketApi
//   var stage WebSocketStage
//
//
//   key := apigwv2.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &RateLimitedApiKeyProps{
//   	CustomerId: jsii.String("test-customer"),
//   	ApiStages: []UsagePlanPerApiStage{
//   		&UsagePlanPerApiStage{
//   			Api: api,
//   			Stage: stage,
//   		},
//   	},
//   	Quota: &QuotaSettings{
//   		Limit: jsii.Number(10000),
//   		Period: apigwv2.Period_MONTH,
//   	},
//   	Throttle: &ThrottleSettings{
//   		RateLimit: jsii.Number(100),
//   		BurstLimit: jsii.Number(200),
//   	},
//   })
//
type QuotaSettings struct {
	// The maximum number of requests that users can make within the specified time period.
	// Default: none.
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// For the initial time period, the number of requests to subtract from the specified limit.
	// Default: none.
	//
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period for which the maximum limit of requests applies.
	// Default: none.
	//
	Period Period `field:"optional" json:"period" yaml:"period"`
}

