package awsapigateway


// The `MethodSetting` property type configures settings for all methods in a stage.
//
// The `MethodSettings` property of the `AWS::ApiGateway::Stage` resource contains a list of `MethodSetting` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodSettingProperty := &methodSettingProperty{
//   	cacheDataEncrypted: jsii.Boolean(false),
//   	cacheTtlInSeconds: jsii.Number(123),
//   	cachingEnabled: jsii.Boolean(false),
//   	dataTraceEnabled: jsii.Boolean(false),
//   	httpMethod: jsii.String("httpMethod"),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	metricsEnabled: jsii.Boolean(false),
//   	resourcePath: jsii.String("resourcePath"),
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   }
//
type CfnStage_MethodSettingProperty struct {
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Indicates whether responses are cached and returned for requests.
	//
	// You must enable a cache cluster on the stage to cache responses.
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Indicates whether data trace logging is enabled for methods in the stage.
	//
	// API Gateway pushes these logs to Amazon CloudWatch Logs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// The HTTP method.
	//
	// To apply settings to multiple resources and methods, specify an asterisk ( `*` ) for the `HttpMethod` and `/*` for the `ResourcePath` . This parameter is required when you specify a `MethodSetting` .
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// The logging level for this method.
	//
	// For valid values, see the `loggingLevel` property of the [Stage](https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the *Amazon API Gateway API Reference* .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// The resource path for this method.
	//
	// Forward slashes ( `/` ) are encoded as `~1` and the initial slash must include a forward slash. For example, the path value `/resource/subresource` must be encoded as `/~1resource~1subresource` . To specify the root path, use only a slash ( `/` ). To apply settings to multiple resources and methods, specify an asterisk ( `*` ) for the `HttpMethod` and `/*` for the `ResourcePath` . This parameter is required when you specify a `MethodSetting` .
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// The number of burst requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account .
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// The number of steady-state requests per second that API Gateway permits across all APIs, stages, and methods in your AWS account .
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

