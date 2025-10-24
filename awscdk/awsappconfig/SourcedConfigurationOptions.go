package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Options for SourcedConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationSource ConfigurationSource
//   var deploymentStrategy DeploymentStrategy
//   var environment Environment
//   var key Key
//   var roleRef IRoleRef
//   var validator IValidator
//
//   sourcedConfigurationOptions := &SourcedConfigurationOptions{
//   	Location: configurationSource,
//
//   	// the properties below are optional
//   	DeletionProtectionCheck: awscdk.Aws_appconfig.DeletionProtectionCheck_ACCOUNT_DEFAULT,
//   	DeploymentKey: key,
//   	DeploymentStrategy: deploymentStrategy,
//   	DeployTo: []IEnvironment{
//   		environment,
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RetrievalRole: roleRef,
//   	Type: awscdk.*Aws_appconfig.ConfigurationType_FREEFORM,
//   	Validators: []IValidator{
//   		validator,
//   	},
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
type SourcedConfigurationOptions struct {
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
	// The location where the configuration is stored.
	Location ConfigurationSource `field:"required" json:"location" yaml:"location"`
	// The IAM role to retrieve the configuration.
	// Default: - A role is generated.
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

