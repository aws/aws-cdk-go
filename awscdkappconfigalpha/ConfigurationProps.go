package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for the Configuration construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import appconfig_alpha "github.com/aws/aws-cdk-go/awscdkappconfigalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var application application
//   var deploymentStrategy deploymentStrategy
//   var environment environment
//   var key key
//   var validator iValidator
//
//   configurationProps := &ConfigurationProps{
//   	Application: application,
//
//   	// the properties below are optional
//   	DeploymentKey: key,
//   	DeploymentStrategy: deploymentStrategy,
//   	DeployTo: []iEnvironment{
//   		environment,
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Type: appconfig_alpha.ConfigurationType_FREEFORM,
//   	Validators: []*iValidator{
//   		validator,
//   	},
//   }
//
// Experimental.
type ConfigurationProps struct {
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
	// The application associated with the configuration.
	// Experimental.
	Application IApplication `field:"required" json:"application" yaml:"application"`
}

