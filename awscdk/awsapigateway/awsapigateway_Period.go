package awsapigateway


// Time period for which quota settings apply.
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
type Period string

const (
	Period_DAY Period = "DAY"
	Period_WEEK Period = "WEEK"
	Period_MONTH Period = "MONTH"
)

