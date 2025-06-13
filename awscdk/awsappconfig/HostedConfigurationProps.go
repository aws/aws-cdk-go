package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for HostedConfiguration.
//
// Example:
//   var application application
//   var fn function
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Validators: []iValidator{
//   		appconfig.JsonSchemaValidator_FromFile(jsii.String("schema.json")),
//   		appconfig.LambdaValidator_FromFunction(fn),
//   	},
//   })
//
type HostedConfigurationProps struct {
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
	// The content of the hosted configuration.
	Content ConfigurationContent `field:"required" json:"content" yaml:"content"`
	// The customer managed key to encrypt hosted configuration.
	// Default: None.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The latest version number of the hosted configuration.
	// Default: - None.
	//
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// The version label of the hosted configuration.
	// Default: - None.
	//
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

