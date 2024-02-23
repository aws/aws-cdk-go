package awscdkappconfigalpha


// The configuration type.
// Deprecated.
type ConfigurationType string

const (
	// Freeform configuration profile.
	//
	// Allows you to store your data in the AWS AppConfig
	// hosted configuration store or another Systems Manager capability or AWS service that integrates
	// with AWS AppConfig.
	// See: https://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-free-form-configurations-creating.html
	//
	// Deprecated.
	ConfigurationType_FREEFORM ConfigurationType = "FREEFORM"
	// Feature flag configuration profile.
	//
	// This configuration stores its data
	// in the AWS AppConfig hosted configuration store and the URI is simply hosted.
	// Deprecated.
	ConfigurationType_FEATURE_FLAGS ConfigurationType = "FEATURE_FLAGS"
)

