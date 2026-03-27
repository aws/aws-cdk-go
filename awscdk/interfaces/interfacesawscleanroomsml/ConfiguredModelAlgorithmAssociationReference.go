package interfacesawscleanroomsml


// A reference to a ConfiguredModelAlgorithmAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredModelAlgorithmAssociationReference := &ConfiguredModelAlgorithmAssociationReference{
//   	ConfiguredModelAlgorithmAssociationArn: jsii.String("configuredModelAlgorithmAssociationArn"),
//   }
//
type ConfiguredModelAlgorithmAssociationReference struct {
	// The ConfiguredModelAlgorithmAssociationArn of the ConfiguredModelAlgorithmAssociation resource.
	ConfiguredModelAlgorithmAssociationArn *string `field:"required" json:"configuredModelAlgorithmAssociationArn" yaml:"configuredModelAlgorithmAssociationArn"`
}

