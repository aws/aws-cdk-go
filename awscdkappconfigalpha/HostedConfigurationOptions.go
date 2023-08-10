package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

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
//   	ContentType: jsii.String("contentType"),
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
// Experimental.
type HostedConfigurationOptions struct {
	// The deployment key of the configuration.
	// Experimental.
	DeploymentKey awskms.IKey `field:"optional" json:"deploymentKey" yaml:"deploymentKey"`
	// The deployment strategy for the configuration.
	// Experimental.
	DeploymentStrategy IDeploymentStrategy `field:"optional" json:"deploymentStrategy" yaml:"deploymentStrategy"`
	// The list of environments to deploy the configuration to.
	//
	// If this parameter is not specified and there is only one environment
	// associated to the application, then we will deploy to that one. Otherwise,
	// there will be no deployment.
	// Experimental.
	DeployTo *[]IEnvironment `field:"optional" json:"deployTo" yaml:"deployTo"`
	// The description of the configuration.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the configuration.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of configuration.
	// Experimental.
	Type ConfigurationType `field:"optional" json:"type" yaml:"type"`
	// The validators for the configuration.
	// Experimental.
	Validators *[]IValidator `field:"optional" json:"validators" yaml:"validators"`
	// The content of the hosted configuration.
	// Experimental.
	Content ConfigurationContent `field:"required" json:"content" yaml:"content"`
	// The content type of the hosted configuration.
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The latest version number of the hosted configuration.
	// Experimental.
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The version label of the hosted configuration.
	// Experimental.
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

