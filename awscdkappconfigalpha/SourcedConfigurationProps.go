package awscdkappconfigalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   var application application
//   var bucket bucket
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
// Experimental.
type SourcedConfigurationProps struct {
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

