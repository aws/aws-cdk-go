package awsapigateway


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &cfnDeploymentProps{
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	deploymentCanarySettings: &deploymentCanarySettingsProperty{
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	description: jsii.String("description"),
//   	stageDescription: &stageDescriptionProperty{
//   		accessLogSetting: &accessLogSettingProperty{
//   			destinationArn: jsii.String("destinationArn"),
//   			format: jsii.String("format"),
//   		},
//   		cacheClusterEnabled: jsii.Boolean(false),
//   		cacheClusterSize: jsii.String("cacheClusterSize"),
//   		cacheDataEncrypted: jsii.Boolean(false),
//   		cacheTtlInSeconds: jsii.Number(123),
//   		cachingEnabled: jsii.Boolean(false),
//   		canarySetting: &canarySettingProperty{
//   			percentTraffic: jsii.Number(123),
//   			stageVariableOverrides: map[string]*string{
//   				"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   			},
//   			useStageCache: jsii.Boolean(false),
//   		},
//   		clientCertificateId: jsii.String("clientCertificateId"),
//   		dataTraceEnabled: jsii.Boolean(false),
//   		description: jsii.String("description"),
//   		documentationVersion: jsii.String("documentationVersion"),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		methodSettings: []interface{}{
//   			&methodSettingProperty{
//   				cacheDataEncrypted: jsii.Boolean(false),
//   				cacheTtlInSeconds: jsii.Number(123),
//   				cachingEnabled: jsii.Boolean(false),
//   				dataTraceEnabled: jsii.Boolean(false),
//   				httpMethod: jsii.String("httpMethod"),
//   				loggingLevel: jsii.String("loggingLevel"),
//   				metricsEnabled: jsii.Boolean(false),
//   				resourcePath: jsii.String("resourcePath"),
//   				throttlingBurstLimit: jsii.Number(123),
//   				throttlingRateLimit: jsii.Number(123),
//   			},
//   		},
//   		metricsEnabled: jsii.Boolean(false),
//   		tags: []cfnTag{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   		tracingEnabled: jsii.Boolean(false),
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   }
//
type CfnDeploymentProps struct {
	// The ID of the `RestApi` resource to deploy.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Specifies settings for the canary deployment.
	DeploymentCanarySettings interface{} `field:"optional" json:"deploymentCanarySettings" yaml:"deploymentCanarySettings"`
	// A description of the purpose of the API Gateway deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures the stage that API Gateway creates with this deployment.
	StageDescription interface{} `field:"optional" json:"stageDescription" yaml:"stageDescription"`
	// A name for the stage that API Gateway creates with this deployment.
	//
	// Use only alphanumeric characters.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

