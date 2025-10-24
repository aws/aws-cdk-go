package awsapigatewayv2


// Time period for which quota settings apply.
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
type Period string

const (
	// The quota resets every day.
	Period_DAY Period = "DAY"
	// The quota resets every week.
	Period_WEEK Period = "WEEK"
	// The quota resets every month.
	Period_MONTH Period = "MONTH"
)

