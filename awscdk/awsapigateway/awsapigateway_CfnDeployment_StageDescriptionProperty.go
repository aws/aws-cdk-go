package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// `StageDescription` is a property of the [AWS::ApiGateway::Deployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html) resource that configures a deployment stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageDescriptionProperty := &stageDescriptionProperty{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	cacheDataEncrypted: jsii.Boolean(false),
//   	cacheTtlInSeconds: jsii.Number(123),
//   	cachingEnabled: jsii.Boolean(false),
//   	canarySetting: &canarySettingProperty{
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	clientCertificateId: jsii.String("clientCertificateId"),
//   	dataTraceEnabled: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	documentationVersion: jsii.String("documentationVersion"),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	methodSettings: []interface{}{
//   		&methodSettingProperty{
//   			cacheDataEncrypted: jsii.Boolean(false),
//   			cacheTtlInSeconds: jsii.Number(123),
//   			cachingEnabled: jsii.Boolean(false),
//   			dataTraceEnabled: jsii.Boolean(false),
//   			httpMethod: jsii.String("httpMethod"),
//   			loggingLevel: jsii.String("loggingLevel"),
//   			metricsEnabled: jsii.Boolean(false),
//   			resourcePath: jsii.String("resourcePath"),
//   			throttlingBurstLimit: jsii.Number(123),
//   			throttlingRateLimit: jsii.Number(123),
//   		},
//   	},
//   	metricsEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnDeployment_StageDescriptionProperty struct {
	// Specifies settings for logging access in this stage.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The size of the stage's cache cluster.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Indicates whether the cached responses are encrypted.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Indicates whether responses are cached and returned for requests.
	//
	// You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide* .
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies settings for the canary deployment in this stage.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// Indicates whether data trace logging is enabled for methods in the stage.
	//
	// API Gateway pushes these logs to Amazon CloudWatch Logs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// A description of the purpose of the stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version identifier of the API documentation snapshot.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// The logging level for this method.
	//
	// For valid values, see the `loggingLevel` property of the [Stage](https://docs.aws.amazon.com/apigateway/api-reference/resource/stage/#loggingLevel) resource in the *Amazon API Gateway API Reference* .
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Configures settings for all of the stage's methods.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// An array of arbitrary tags (key-value pairs) to associate with the stage.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The target request burst rate limit.
	//
	// This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// The target request steady-state rate limit.
	//
	// For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide* .
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
	// Specifies whether active tracing with X-ray is enabled for this stage.
	//
	// For more information, see [Trace API Gateway API Execution with AWS X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide* .
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map that defines the stage variables.
	//
	// Variable names must consist of alphanumeric characters, and the values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

