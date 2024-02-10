package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnStage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStageProps := &CfnStageProps{
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	AccessLogSetting: &AccessLogSettingProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	CacheClusterEnabled: jsii.Boolean(false),
//   	CacheClusterSize: jsii.String("cacheClusterSize"),
//   	CanarySetting: &CanarySettingProperty{
//   		DeploymentId: jsii.String("deploymentId"),
//   		PercentTraffic: jsii.Number(123),
//   		StageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		UseStageCache: jsii.Boolean(false),
//   	},
//   	ClientCertificateId: jsii.String("clientCertificateId"),
//   	DeploymentId: jsii.String("deploymentId"),
//   	Description: jsii.String("description"),
//   	DocumentationVersion: jsii.String("documentationVersion"),
//   	MethodSettings: []interface{}{
//   		&MethodSettingProperty{
//   			CacheDataEncrypted: jsii.Boolean(false),
//   			CacheTtlInSeconds: jsii.Number(123),
//   			CachingEnabled: jsii.Boolean(false),
//   			DataTraceEnabled: jsii.Boolean(false),
//   			HttpMethod: jsii.String("httpMethod"),
//   			LoggingLevel: jsii.String("loggingLevel"),
//   			MetricsEnabled: jsii.Boolean(false),
//   			ResourcePath: jsii.String("resourcePath"),
//   			ThrottlingBurstLimit: jsii.Number(123),
//   			ThrottlingRateLimit: jsii.Number(123),
//   		},
//   	},
//   	StageName: jsii.String("stageName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TracingEnabled: jsii.Boolean(false),
//   	Variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html
//
type CfnStageProps struct {
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Access log settings, including the access log format and access log destination ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-accesslogsetting
	//
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Specifies whether a cache cluster is enabled for the stage.
	//
	// To activate a method-level cache, set `CachingEnabled` to `true` for a method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclusterenabled
	//
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache capacity in GB.
	//
	// For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-cacheclustersize
	//
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Settings for the canary deployment in this stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-canarysetting
	//
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The identifier of a client certificate for an API stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-clientcertificateid
	//
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The identifier of the Deployment that the stage points to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-deploymentid
	//
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The stage's description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the associated API documentation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-documentationversion
	//
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// A map that defines the method settings for a Stage resource.
	//
	// Keys (designated as `/{method_setting_key` below) are method paths defined as `{resource_path}/{http_method}` for an individual method override, or `/\* /\*` for overriding all methods in the stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-methodsettings
	//
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway.
	//
	// Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-stagename
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether active tracing with X-ray is enabled for the Stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-tracingenabled
	//
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
	//
	// Variable names are limited to alphanumeric characters. Values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html#cfn-apigateway-stage-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

