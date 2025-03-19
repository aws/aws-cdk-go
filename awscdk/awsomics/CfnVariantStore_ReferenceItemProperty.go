package awsomics


// The read set's genome reference ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceItemProperty := &ReferenceItemProperty{
//   	ReferenceArn: jsii.String("referenceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-variantstore-referenceitem.html
//
type CfnVariantStore_ReferenceItemProperty struct {
	// The reference's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-variantstore-referenceitem.html#cfn-omics-variantstore-referenceitem-referencearn
	//
	ReferenceArn *string `field:"required" json:"referenceArn" yaml:"referenceArn"`
}

