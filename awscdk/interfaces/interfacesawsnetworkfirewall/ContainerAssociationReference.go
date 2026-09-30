package interfacesawsnetworkfirewall


// A reference to a ContainerAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerAssociationReference := &ContainerAssociationReference{
//   	ContainerAssociationArn: jsii.String("containerAssociationArn"),
//   }
//
type ContainerAssociationReference struct {
	// The ContainerAssociationArn of the ContainerAssociation resource.
	ContainerAssociationArn *string `field:"required" json:"containerAssociationArn" yaml:"containerAssociationArn"`
}

