package previewawsomicsmixins


// A genome reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceItemProperty := &ReferenceItemProperty{
//   	ReferenceArn: jsii.String("referenceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-referenceitem.html
//
type CfnAnnotationStorePropsMixin_ReferenceItemProperty struct {
	// The reference's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-annotationstore-referenceitem.html#cfn-omics-annotationstore-referenceitem-referencearn
	//
	ReferenceArn *string `field:"optional" json:"referenceArn" yaml:"referenceArn"`
}

