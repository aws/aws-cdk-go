package awsapigateway


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	DeploymentCanarySettings: &DeploymentCanarySettingsProperty{
//   		PercentTraffic: jsii.Number(123),
//   		StageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		UseStageCache: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	StageDescription: &StageDescriptionProperty{
//   		AccessLogSetting: &AccessLogSettingProperty{
//   			DestinationArn: jsii.String("destinationArn"),
//   			Format: jsii.String("format"),
//   		},
//   		CacheClusterEnabled: jsii.Boolean(false),
//   		CacheClusterSize: jsii.String("cacheClusterSize"),
//   		CacheDataEncrypted: jsii.Boolean(false),
//   		CacheTtlInSeconds: jsii.Number(123),
//   		CachingEnabled: jsii.Boolean(false),
//   		CanarySetting: &CanarySettingProperty{
//   			PercentTraffic: jsii.Number(123),
//   			StageVariableOverrides: map[string]*string{
//   				"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   			},
//   			UseStageCache: jsii.Boolean(false),
//   		},
//   		ClientCertificateId: jsii.String("clientCertificateId"),
//   		DataTraceEnabled: jsii.Boolean(false),
//   		Description: jsii.String("description"),
//   		DocumentationVersion: jsii.String("documentationVersion"),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		MethodSettings: []interface{}{
//   			&MethodSettingProperty{
//   				CacheDataEncrypted: jsii.Boolean(false),
//   				CacheTtlInSeconds: jsii.Number(123),
//   				CachingEnabled: jsii.Boolean(false),
//   				DataTraceEnabled: jsii.Boolean(false),
//   				HttpMethod: jsii.String("httpMethod"),
//   				LoggingLevel: jsii.String("loggingLevel"),
//   				MetricsEnabled: jsii.Boolean(false),
//   				ResourcePath: jsii.String("resourcePath"),
//   				ThrottlingBurstLimit: jsii.Number(123),
//   				ThrottlingRateLimit: jsii.Number(123),
//   			},
//   		},
//   		MetricsEnabled: jsii.Boolean(false),
//   		Tags: []cfnTag{
//   			&cfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   		TracingEnabled: jsii.Boolean(false),
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	StageName: jsii.String("stageName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
//
type CfnDeploymentProps struct {
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The input configuration for a canary deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-deploymentcanarysettings
	//
	DeploymentCanarySettings interface{} `field:"optional" json:"deploymentCanarySettings" yaml:"deploymentCanarySettings"`
	// The description for the Deployment resource to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The description of the Stage resource for the Deployment resource to create.
	//
	// To specify a stage description, you must also provide a stage name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagedescription
	//
	StageDescription interface{} `field:"optional" json:"stageDescription" yaml:"stageDescription"`
	// The name of the Stage resource for the Deployment resource to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagename
	//
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

