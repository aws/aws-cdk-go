package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customInstructionProperty := &CustomInstructionProperty{
//   	InlineCustomInstruction: &InlineCustomInstructionProperty{
//   		InstructionText: jsii.String("instructionText"),
//
//   		// the properties below are optional
//   		UploadedDocumentMetadata: &UploadedDocumentMetadataProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-custominstruction.html
//
type CfnDataSet_CustomInstructionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-custominstruction.html#cfn-quicksight-dataset-custominstruction-inlinecustominstruction
	//
	InlineCustomInstruction interface{} `field:"optional" json:"inlineCustomInstruction" yaml:"inlineCustomInstruction"`
}

