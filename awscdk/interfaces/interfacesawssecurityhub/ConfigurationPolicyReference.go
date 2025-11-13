package interfacesawssecurityhub


// A reference to a ConfigurationPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationPolicyReference := &ConfigurationPolicyReference{
//   	ConfigurationPolicyArn: jsii.String("configurationPolicyArn"),
//   }
//
type ConfigurationPolicyReference struct {
	// The Arn of the ConfigurationPolicy resource.
	ConfigurationPolicyArn *string `field:"required" json:"configurationPolicyArn" yaml:"configurationPolicyArn"`
}

