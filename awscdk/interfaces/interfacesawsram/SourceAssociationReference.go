package interfacesawsram


// A reference to a SourceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceAssociationReference := &SourceAssociationReference{
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   	SourceId: jsii.String("sourceId"),
//   }
//
type SourceAssociationReference struct {
	// The ResourceShareArn of the SourceAssociation resource.
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
	// The SourceId of the SourceAssociation resource.
	SourceId *string `field:"required" json:"sourceId" yaml:"sourceId"`
}

