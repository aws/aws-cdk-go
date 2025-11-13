package interfacesawsomics


// A reference to a SequenceStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sequenceStoreReference := &SequenceStoreReference{
//   	SequenceStoreArn: jsii.String("sequenceStoreArn"),
//   	SequenceStoreId: jsii.String("sequenceStoreId"),
//   }
//
type SequenceStoreReference struct {
	// The ARN of the SequenceStore resource.
	SequenceStoreArn *string `field:"required" json:"sequenceStoreArn" yaml:"sequenceStoreArn"`
	// The SequenceStoreId of the SequenceStore resource.
	SequenceStoreId *string `field:"required" json:"sequenceStoreId" yaml:"sequenceStoreId"`
}

