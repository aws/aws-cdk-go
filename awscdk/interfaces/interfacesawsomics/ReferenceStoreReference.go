package interfacesawsomics


// A reference to a ReferenceStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceStoreReference := &ReferenceStoreReference{
//   	ReferenceStoreArn: jsii.String("referenceStoreArn"),
//   	ReferenceStoreId: jsii.String("referenceStoreId"),
//   }
//
type ReferenceStoreReference struct {
	// The ARN of the ReferenceStore resource.
	ReferenceStoreArn *string `field:"required" json:"referenceStoreArn" yaml:"referenceStoreArn"`
	// The ReferenceStoreId of the ReferenceStore resource.
	ReferenceStoreId *string `field:"required" json:"referenceStoreId" yaml:"referenceStoreId"`
}

