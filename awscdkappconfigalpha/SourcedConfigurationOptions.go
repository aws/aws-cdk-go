package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationSource configurationSource
//   var deploymentStrategy deploymentStrategy
//   var environment environment
//   var key key
//   var role role
//   var validator iValidator
//
//   sourcedConfigurationOptions := &SourcedConfigurationOptions{
//   	Location: configurationSource,
//
//   	// the properties below are optional
//   	DeploymentKey: key,
//   	DeploymentStrategy: deploymentStrategy,
//   	DeployTo: []iEnvironment{
//   		environment,
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RetrievalRole: role,
//   	Type: appconfig_alpha.ConfigurationType_FREEFORM,
//   	Validators: []*iValidator{
//   		validator,
//   	},
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
// Experimental.
type SourcedConfigurationOptions struct {
	// The deployment key of the configuration.
	// Default: - None.
	//
	// Experimental.
	DeploymentKey awskms.IKey `field:"optional" json:"deploymentKey" yaml:"deploymentKey"`
	// The deployment strategy for the configuration.
	// Default: - A deployment strategy with the rollout strategy set to
	// RolloutStrategy.CANARY_10_PERCENT_20_MINUTES
	//
	// Experimental.
	DeploymentStrategy IDeploymentStrategy `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
	// The list of environments to deploy the configuration to.
	//
	// If this parameter is not specified, then there will be no
	// deployment.
	// Default: - None.
	//
	// Experimental.
	DeployTo *[]IEnvironment `field:"optional" json:"deployTo" yaml:"deployTo"`
	// The description of the configuration.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the configuration.
	// Default: - A name is generated.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of configuration.
	// Default: ConfigurationType.FREEFORM
	//
	// Experimental.
	Type ConfigurationType `field:"optional" json:"type" yaml:"type"`
	// The validators for the configuration.
	// Default: - No validators.
	//
	// Experimental.
	Validators *[]IValidator `field:"optional" json:"validators" yaml:"validators"`
	// The location where the configuration is stored.
	// Experimental.
	Location ConfigurationSource `field:"required" json:"location" yaml:"location"`
	// The IAM role to retrieve the configuration.
	// Default: - A role is generated.
	//
	// Experimental.
	RetrievalRole awsiam.IRole `field:"optional" json:"retrievalRole" yaml:"retrievalRole"`
	// The version number of the sourced configuration to deploy.
	//
	// If this is not specified,
	// then there will be no deployment.
	// Default: - None.
	//
	// Experimental.
	VersionNumber *string `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

