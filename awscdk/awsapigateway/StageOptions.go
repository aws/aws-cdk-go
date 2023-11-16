package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   logGroup := logs.NewLogGroup(this, jsii.String("ApiGatewayAccessLogs"))
//   api := apigateway.NewRestApi(this, jsii.String("books"), &RestApiProps{
//   	DeployOptions: &StageOptions{
//   		AccessLogDestination: apigateway.NewLogGroupLogDestination(logGroup),
//   		AccessLogFormat: apigateway.AccessLogFormat_Clf(),
//   	},
//   })
//
type StageOptions struct {
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
	// The CloudWatch Logs log group or Firehose delivery stream where to write access logs.
	// Default: - No destination.
	//
	AccessLogDestination IAccessLogDestination `field:"optional" json:"accessLogDestination" yaml:"accessLogDestination"`
	// A single line format of access logs of data, as specified by selected $content variables.
	//
	// The format must include either `AccessLogFormat.contextRequestId()`
	// or `AccessLogFormat.contextExtendedRequestId()`.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference
	//
	// Default: - Common Log Format.
	//
	AccessLogFormat AccessLogFormat `field:"optional" json:"accessLogFormat" yaml:"accessLogFormat"`
	// Indicates whether cache clustering is enabled for the stage.
	// Default: - Disabled for the stage.
	//
	CacheClusterEnabled *bool `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache cluster size.
	// Default: 0.5
	//
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	// Default: - None.
	//
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// A description of the purpose of the stage.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version identifier of the API documentation snapshot.
	// Default: - No documentation version.
	//
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Method deployment options for specific resources/methods.
	//
	// These will
	// override common options defined in `StageOptions#methodOptions`.
	// Default: - Common options will be used.
	//
	MethodOptions *map[string]*MethodDeploymentOptions `field:"optional" json:"methodOptions" yaml:"methodOptions"`
	// The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
	// Default: - "prod".
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// Specifies whether Amazon X-Ray tracing is enabled for this method.
	// Default: false.
	//
	TracingEnabled *bool `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map that defines the stage variables.
	//
	// Variable names must consist of
	// alphanumeric characters, and the values must match the following regular
	// expression: [A-Za-z0-9-._~:/?#&=,]+.
	// Default: - No stage variables.
	//
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

