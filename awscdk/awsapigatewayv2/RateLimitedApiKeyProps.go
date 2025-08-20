package awsapigatewayv2


// RateLimitedApiKey properties.
//
// Example:
//   var api webSocketApi
//   var stage webSocketStage
//
//
//   key := apigwv2.NewRateLimitedApiKey(this, jsii.String("rate-limited-api-key"), &RateLimitedApiKeyProps{
//   	CustomerId: jsii.String("test-customer"),
//   	ApiStages: []usagePlanPerApiStage{
//   		&usagePlanPerApiStage{
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
type RateLimitedApiKeyProps struct {
	// A name for the API key.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.
	// Default: automatically generated name.
	//
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// A description of the purpose of the API key.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value of the API key.
	//
	// Must be at least 20 characters long.
	// Default: none.
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
	// An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.
	// Default: none.
	//
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// Indicates whether the API key can be used by clients.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether the key identifier is distinct from the created API key value.
	// Default: false.
	//
	GenerateDistinctId *bool `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// API Stages to be associated with the RateLimitedApiKey.
	// Default: none.
	//
	ApiStages *[]*UsagePlanPerApiStage `field:"optional" json:"apiStages" yaml:"apiStages"`
	// Number of requests clients can make in a given time period.
	// Default: none.
	//
	Quota *QuotaSettings `field:"optional" json:"quota" yaml:"quota"`
	// Overall throttle settings for the API.
	// Default: none.
	//
	Throttle *ThrottleSettings `field:"optional" json:"throttle" yaml:"throttle"`
}

