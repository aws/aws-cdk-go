package awsappconfig


// The configuration type.
//
// Example:
//   var application Application
//   var kmsKey Key
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
//   	KmsKey: KmsKey,
//   })
//
type ConfigurationType string

const (
	// Freeform configuration profile.
	//
	// Allows you to store your data in the AWS AppConfig
	// hosted configuration store or another Systems Manager capability or AWS service that integrates
	// with AWS AppConfig.
	// See: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-free-form-configurations-creating.html
	//
	ConfigurationType_FREEFORM ConfigurationType = "FREEFORM"
	// Feature flag configuration profile.
	//
	// This configuration stores its data
	// in the AWS AppConfig hosted configuration store and the URI is simply hosted.
	ConfigurationType_FEATURE_FLAGS ConfigurationType = "FEATURE_FLAGS"
)

