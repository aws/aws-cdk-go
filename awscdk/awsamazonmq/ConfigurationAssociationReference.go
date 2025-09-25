package awsamazonmq


// A reference to a ConfigurationAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configurationAssociationReference := &ConfigurationAssociationReference{
//   	ConfigurationAssociationId: jsii.String("configurationAssociationId"),
//   }
//
type ConfigurationAssociationReference struct {
	// The Id of the ConfigurationAssociation resource.
	ConfigurationAssociationId *string `field:"required" json:"configurationAssociationId" yaml:"configurationAssociationId"`
}

