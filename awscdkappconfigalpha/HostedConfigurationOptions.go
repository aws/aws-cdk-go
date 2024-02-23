package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Options for HostedConfiguration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationContent configurationContent
//   var deploymentStrategy deploymentStrategy
//   var environment environment
//   var key key
//   var validator iValidator
//
//   hostedConfigurationOptions := &HostedConfigurationOptions{
//   	Content: configurationContent,
//
//   	// the properties below are optional
//   	DeploymentKey: key,
//   	DeploymentStrategy: deploymentStrategy,
//   	DeployTo: []iEnvironment{
//   		environment,
//   	},
//   	Description: jsii.String("description"),
//   	LatestVersionNumber: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Type: appconfig_alpha.ConfigurationType_FREEFORM,
//   	Validators: []*iValidator{
//   		validator,
//   	},
//   	VersionLabel: jsii.String("versionLabel"),
//   }
//
// Deprecated.
type HostedConfigurationOptions struct {
	// The deployment key of the configuration.
	// Default: - None.
	//
	// Deprecated.
	DeploymentKey awskms.IKey `field:"optional" json:"deploymentKey" yaml:"deploymentKey"`
	// The deployment strategy for the configuration.
	// Default: - A deployment strategy with the rollout strategy set to
	// RolloutStrategy.CANARY_10_PERCENT_20_MINUTES
	//
	// Deprecated.
	DeploymentStrategy IDeploymentStrategy `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
	// The list of environments to deploy the configuration to.
	//
	// If this parameter is not specified, then there will be no
	// deployment.
	// Default: - None.
	//
	// Deprecated.
	DeployTo *[]IEnvironment `field:"optional" json:"deployTo" yaml:"deployTo"`
	// The description of the configuration.
	// Default: - No description.
	//
	// Deprecated.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the configuration.
	// Default: - A name is generated.
	//
	// Deprecated.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of configuration.
	// Default: ConfigurationType.FREEFORM
	//
	// Deprecated.
	Type ConfigurationType `field:"optional" json:"type" yaml:"type"`
	// The validators for the configuration.
	// Default: - No validators.
	//
	// Deprecated.
	Validators *[]IValidator `field:"optional" json:"validators" yaml:"validators"`
	// The content of the hosted configuration.
	// Deprecated.
	Content ConfigurationContent `field:"required" json:"content" yaml:"content"`
	// The latest version number of the hosted configuration.
	// Default: - None.
	//
	// Deprecated.
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The version label of the hosted configuration.
	// Default: - None.
	//
	// Deprecated.
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

