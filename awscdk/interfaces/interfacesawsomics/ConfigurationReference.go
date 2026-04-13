package interfacesawsomics


// A reference to a Configuration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationReference := &ConfigurationReference{
//   	ConfigurationArn: jsii.String("configurationArn"),
//   	ConfigurationName: jsii.String("configurationName"),
//   }
//
type ConfigurationReference struct {
	// The ARN of the Configuration resource.
	ConfigurationArn *string `field:"required" json:"configurationArn" yaml:"configurationArn"`
	// The Name of the Configuration resource.
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
}

