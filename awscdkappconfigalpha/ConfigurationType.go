package awscdkappconfigalpha


// The configuration type.
//
// Example:
//   var application application
//
//
//   appconfig.NewHostedConfiguration(this, jsii.String("MyHostedConfiguration"), &HostedConfigurationProps{
//   	Application: Application,
//   	Content: appconfig.ConfigurationContent_FromInlineText(jsii.String("This is my configuration content.")),
//   	Type: appconfig.ConfigurationType_FEATURE_FLAGS,
//   })
//
// Experimental.
type ConfigurationType string

const (
	// Experimental.
	ConfigurationType_FREEFORM ConfigurationType = "FREEFORM"
	// Experimental.
	ConfigurationType_FEATURE_FLAGS ConfigurationType = "FEATURE_FLAGS"
)

