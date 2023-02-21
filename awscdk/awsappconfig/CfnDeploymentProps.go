package awsappconfig


// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   	ConfigurationVersion: jsii.String("configurationVersion"),
//   	DeploymentStrategyId: jsii.String("deploymentStrategyId"),
//   	EnvironmentId: jsii.String("environmentId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	Tags: []tagsProperty{
//   		&tagsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeploymentProps struct {
	// The application ID.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The configuration profile ID.
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The configuration version to deploy.
	//
	// If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label.
	ConfigurationVersion *string `field:"required" json:"configurationVersion" yaml:"configurationVersion"`
	// The deployment strategy ID.
	DeploymentStrategyId *string `field:"required" json:"deploymentStrategyId" yaml:"deploymentStrategyId"`
	// The environment ID.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// A description of the deployment.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::AppConfig::Deployment.KmsKeyIdentifier`.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Metadata to assign to the deployment.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*CfnDeployment_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

