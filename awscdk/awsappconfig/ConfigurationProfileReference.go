package awsappconfig


// A reference to a ConfigurationProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationProfileReference := &ConfigurationProfileReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   }
//
type ConfigurationProfileReference struct {
	// The ApplicationId of the ConfigurationProfile resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ConfigurationProfileId of the ConfigurationProfile resource.
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
}

