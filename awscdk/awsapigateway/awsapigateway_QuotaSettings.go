package awsapigateway


// Specifies the maximum number of requests that clients can make to API Gateway APIs.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &rateLimitedApiKeyProps{
//   	customerId: jsii.String("hello-customer"),
//   	stages: []iStage{
//   		api.deploymentStage,
//   	},
//   	quota: &quotaSettings{
//   		limit: jsii.Number(10000),
//   		period: apigateway.period_MONTH,
//   	},
//   })
//
type QuotaSettings struct {
	// The maximum number of requests that users can make within the specified time period.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// For the initial time period, the number of requests to subtract from the specified limit.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// The time period for which the maximum limit of requests applies.
	Period Period `field:"optional" json:"period" yaml:"period"`
}

