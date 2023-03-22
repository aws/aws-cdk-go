package awsapigateway


// Time period for which quota settings apply.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &rateLimitedApiKeyProps{
//   	customerId: jsii.String("hello-customer"),
//   	resources: []iRestApi{
//   		api,
//   	},
//   	quota: &quotaSettings{
//   		limit: jsii.Number(10000),
//   		period: apigateway.period_MONTH,
//   	},
//   })
//
// Experimental.
type Period string

const (
	// Experimental.
	Period_DAY Period = "DAY"
	// Experimental.
	Period_WEEK Period = "WEEK"
	// Experimental.
	Period_MONTH Period = "MONTH"
)

