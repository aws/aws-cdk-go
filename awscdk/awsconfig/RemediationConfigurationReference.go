package awsconfig


// A reference to a RemediationConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remediationConfigurationReference := &RemediationConfigurationReference{
//   	RemediationConfigurationId: jsii.String("remediationConfigurationId"),
//   }
//
type RemediationConfigurationReference struct {
	// The Id of the RemediationConfiguration resource.
	RemediationConfigurationId *string `field:"required" json:"remediationConfigurationId" yaml:"remediationConfigurationId"`
}

