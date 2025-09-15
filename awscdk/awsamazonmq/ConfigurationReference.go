package awsamazonmq


// A reference to a Configuration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationReference := &ConfigurationReference{
//   	ConfigurationArn: jsii.String("configurationArn"),
//   	ConfigurationId: jsii.String("configurationId"),
//   }
//
type ConfigurationReference struct {
	// The ARN of the Configuration resource.
	ConfigurationArn *string `field:"required" json:"configurationArn" yaml:"configurationArn"`
	// The Id of the Configuration resource.
	ConfigurationId *string `field:"required" json:"configurationId" yaml:"configurationId"`
}

