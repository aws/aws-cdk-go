package interfacesawsnetworkmanager


// A reference to a ContextStoreAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contextStoreAssociationReference := &ContextStoreAssociationReference{
//   	ContextStoreId: jsii.String("contextStoreId"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type ContextStoreAssociationReference struct {
	// The ContextStoreId of the ContextStoreAssociation resource.
	ContextStoreId *string `field:"required" json:"contextStoreId" yaml:"contextStoreId"`
	// The ResourceArn of the ContextStoreAssociation resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

