package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStageProps := &cfnStageProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	clientCertificateId: jsii.String("clientCertificateId"),
//   	deploymentId: jsii.String("deploymentId"),
//   	description: jsii.String("description"),
//   	documentationVersion: jsii.String("documentationVersion"),
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
//   	stageName: jsii.String("stageName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnStageProps struct {
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Access log settings, including the access log format and access log destination ARN.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache capacity in GB.
	//
	// For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) .
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Settings for the canary deployment in this stage.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The identifier of the Deployment that the stage points to.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The stage's description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the associated API documentation.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// A map that defines the method settings for a Stage resource.
	//
	// Keys (designated as `/{method_setting_key` below) are method paths defined as `{resource_path}/{http_method}` for an individual method override, or `/\* /\*` for overriding all methods in the stage.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway.
	//
	// Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
	//
	// Variable names are limited to alphanumeric characters. Values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

