package interfacesawsconfig


// A reference to a RemediationConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   remediationConfigurationReference := &RemediationConfigurationReference{
//   	ConfigRuleName: jsii.String("configRuleName"),
//   }
//
type RemediationConfigurationReference struct {
	// The ConfigRuleName of the RemediationConfiguration resource.
	ConfigRuleName *string `field:"required" json:"configRuleName" yaml:"configRuleName"`
}

