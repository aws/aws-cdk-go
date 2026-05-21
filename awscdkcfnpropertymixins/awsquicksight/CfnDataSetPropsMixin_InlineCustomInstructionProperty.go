package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inlineCustomInstructionProperty := &InlineCustomInstructionProperty{
//   	InstructionText: jsii.String("instructionText"),
//   	UploadedDocumentMetadata: &UploadedDocumentMetadataProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inlinecustominstruction.html
//
type CfnDataSetPropsMixin_InlineCustomInstructionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inlinecustominstruction.html#cfn-quicksight-dataset-inlinecustominstruction-instructiontext
	//
	InstructionText *string `field:"optional" json:"instructionText" yaml:"instructionText"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-inlinecustominstruction.html#cfn-quicksight-dataset-inlinecustominstruction-uploadeddocumentmetadata
	//
	UploadedDocumentMetadata interface{} `field:"optional" json:"uploadedDocumentMetadata" yaml:"uploadedDocumentMetadata"`
}

