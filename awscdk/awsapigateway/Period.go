package awsapigateway


// Time period for which quota settings apply.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &RateLimitedApiKeyProps{
//   	CustomerId: jsii.String("hello-customer"),
//   	Stages: []iStage{
//   		api.DeploymentStage,
//   	},
//   	Quota: &QuotaSettings{
//   		Limit: jsii.Number(10000),
//   		Period: apigateway.Period_MONTH,
//   	},
//   })
//
type Period string

const (
	Period_DAY Period = "DAY"
	Period_WEEK Period = "WEEK"
	Period_MONTH Period = "MONTH"
)

