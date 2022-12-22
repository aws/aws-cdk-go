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
	// The ID of the `RestApi` resource that you're deploying with this stage.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Specifies settings for logging access in this stage.
	AccessLogSetting interface{} `field:"optional" json:"accessLogSetting" yaml:"accessLogSetting"`
	// Indicates whether cache clustering is enabled for the stage.
	CacheClusterEnabled interface{} `field:"optional" json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// The stage's cache cluster size.
	CacheClusterSize *string `field:"optional" json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// Specifies settings for the canary deployment in this stage.
	CanarySetting interface{} `field:"optional" json:"canarySetting" yaml:"canarySetting"`
	// The ID of the client certificate that API Gateway uses to call your integration endpoints in the stage.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The ID of the deployment that the stage is associated with.
	//
	// This parameter is required to create a stage.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// A description of the stage.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version ID of the API documentation snapshot.
	DocumentationVersion *string `field:"optional" json:"documentationVersion" yaml:"documentationVersion"`
	// Settings for all methods in the stage.
	MethodSettings interface{} `field:"optional" json:"methodSettings" yaml:"methodSettings"`
	// The name of the stage, which API Gateway uses as the first path segment in the invoked Uniform Resource Identifier (URI).
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
	// An array of arbitrary tags (key-value pairs) to associate with the stage.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether active X-Ray tracing is enabled for this stage.
	//
	// For more information, see [Trace API Gateway API Execution with AWS X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide* .
	TracingEnabled interface{} `field:"optional" json:"tracingEnabled" yaml:"tracingEnabled"`
	// A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value.
	//
	// Variable names are limited to alphanumeric characters. Values must match the following regular expression: `[A-Za-z0-9-._~:/?#&=,]+` .
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

