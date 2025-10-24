package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for SourcedConfiguration.
//
// Example:
//   var application Application
//   var bucket Bucket
//
//
//   appconfig.NewSourcedConfiguration(this, jsii.String("MySourcedConfiguration"), &SourcedConfigurationProps{
//   	Application: Application,
//   	Location: appconfig.ConfigurationSource_FromBucket(bucket, jsii.String("path/to/file.json")),
//   	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
//   	Name: jsii.String("MyConfig"),
//   	Description: jsii.String("This is my sourced configuration from CDK."),
//   })
//
type SourcedConfigurationProps struct {
	// A parameter to configure deletion protection.
	//
	// Deletion protection prevents a user from deleting a configuration profile if your application has called
	// either `GetLatestConfiguration` or `GetConfiguration` for the configuration profile during the specified interval.
	// See: https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html
	//
	// Default: DeletionProtectionCheck.ACCOUNT_DEFAULT
	//
	DeletionProtectionCheck DeletionProtectionCheck `field:"optional" json:"deletionProtectionCheck" yaml:"deletionProtectionCheck"`
	// The deployment key of the configuration.
	// Default: - None.
	//
	DeploymentKey awskms.IKey `field:"optional" json:"deploymentKey" yaml:"deploymentKey"`
	// The deployment strategy for the configuration.
	// Default: - A deployment strategy with the rollout strategy set to
	// RolloutStrategy.CANARY_10_PERCENT_20_MINUTES
	//
	DeploymentStrategy IDeploymentStrategy `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
	// The list of environments to deploy the configuration to.
	//
	// If this parameter is not specified, then there will be no
	// deployment created alongside this configuration.
	//
	// Deployments can be added later using the `IEnvironment.addDeployment` or
	// `IEnvironment.addDeployments` methods.
	// Default: - None.
	//
	DeployTo *[]IEnvironment `field:"optional" json:"deployTo" yaml:"deployTo"`
	// The description of the configuration.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the configuration.
	// Default: - A name is generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of configuration.
	// Default: ConfigurationType.FREEFORM
	//
	Type ConfigurationType `field:"optional" json:"type" yaml:"type"`
	// The validators for the configuration.
	// Default: - No validators.
	//
	Validators *[]IValidator `field:"optional" json:"validators" yaml:"validators"`
	// The application associated with the configuration.
	Application IApplication `field:"required" json:"application" yaml:"application"`
	// The location where the configuration is stored.
	Location ConfigurationSource `field:"required" json:"location" yaml:"location"`
	// The IAM role to retrieve the configuration.
	// Default: - Auto generated if location type is not ConfigurationSourceType.CODE_PIPELINE otherwise no role specified.
	//
	RetrievalRole awsiam.IRoleRef `field:"optional" json:"retrievalRole" yaml:"retrievalRole"`
	// The version number of the sourced configuration to deploy.
	//
	// If this is not specified,
	// then there will be no deployment.
	// Default: - None.
	//
	VersionNumber *string `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

