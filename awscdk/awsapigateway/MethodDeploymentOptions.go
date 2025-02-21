package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   api := apigateway.NewRestApi(this, jsii.String("books"))
//   deployment := apigateway.NewDeployment(this, jsii.String("my-deployment"), &DeploymentProps{
//   	Api: Api,
//   })
//   stage := apigateway.NewStage(this, jsii.String("my-stage"), &StageProps{
//   	Deployment: Deployment,
//   	MethodOptions: map[string]methodDeploymentOptions{
//   		"/*/*": &methodDeploymentOptions{
//   			 // This special path applies to all resource paths and all HTTP methods
//   			"throttlingRateLimit": jsii.Number(100),
//   			"throttlingBurstLimit": jsii.Number(200),
//   		},
//   	},
//   })
//
type MethodDeploymentOptions struct {
	// Indicates whether the cached responses are encrypted.
	// Default: false.
	//
	CacheDataEncrypted *bool `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Specifies the time to live (TTL), in seconds, for cached responses.
	//
	// The
	// higher the TTL, the longer the response will be cached.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html
	//
	// Default: Duration.minutes(5)
	//
	CacheTtl awscdk.Duration `field:"optional" json:"cacheTtl" yaml:"cacheTtl"`
	// Specifies whether responses should be cached and returned for requests.
	//
	// A
	// cache cluster must be enabled on the stage for responses to be cached.
	// Default: - Caching is Disabled.
	//
	CachingEnabled *bool `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies whether data trace logging is enabled for this method.
	//
	// When enabled, API gateway will log the full API requests and responses.
	// This can be useful to troubleshoot APIs, but can result in logging sensitive data.
	// We recommend that you don't enable this feature for production APIs.
	// Default: false.
	//
	DataTraceEnabled *bool `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// Specifies the logging level for this method, which effects the log entries pushed to Amazon CloudWatch Logs.
	// Default: - Off.
	//
	LoggingLevel MethodLoggingLevel `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies whether Amazon CloudWatch metrics are enabled for this method.
	// Default: false.
	//
	MetricsEnabled *bool `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Specifies the throttling burst limit.
	//
	// The total rate of all requests in your AWS account is limited to 5,000 requests.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	// Default: - No additional restriction.
	//
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	//
	// The total rate of all requests in your AWS account is limited to 10,000 requests per second (rps).
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html
	//
	// Default: - No additional restriction.
	//
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

