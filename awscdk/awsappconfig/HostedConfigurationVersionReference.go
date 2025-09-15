package awsappconfig


// A reference to a HostedConfigurationVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedConfigurationVersionReference := &HostedConfigurationVersionReference{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
type HostedConfigurationVersionReference struct {
	// The ApplicationId of the HostedConfigurationVersion resource.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The ConfigurationProfileId of the HostedConfigurationVersion resource.
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The VersionNumber of the HostedConfigurationVersion resource.
	VersionNumber *string `field:"required" json:"versionNumber" yaml:"versionNumber"`
}

