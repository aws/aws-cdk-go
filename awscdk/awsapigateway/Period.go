package awsapigateway


// Time period for which quota settings apply.
//
// Example:
//   var api restApi
//
//
//   key := apigateway.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &RateLimitedApiKeyProps{
//   	CustomerId: jsii.String("hello-customer"),
//   	Resources: []iRestApi{
//   		api,
//   	},
//   	Quota: &QuotaSettings{
//   		Limit: jsii.Number(10000),
//   		Period: apigateway.Period_MONTH,
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

