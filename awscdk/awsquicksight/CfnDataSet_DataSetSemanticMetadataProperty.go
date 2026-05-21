package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetSemanticMetadataProperty := &DataSetSemanticMetadataProperty{
//   	CustomInstructions: []interface{}{
//   		&CustomInstructionProperty{
//   			InlineCustomInstruction: &InlineCustomInstructionProperty{
//   				InstructionText: jsii.String("instructionText"),
//
//   				// the properties below are optional
//   				UploadedDocumentMetadata: &UploadedDocumentMetadataProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	Description: &DataSetSemanticDescriptionProperty{
//   		Text: jsii.String("text"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetsemanticmetadata.html
//
type CfnDataSet_DataSetSemanticMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetsemanticmetadata.html#cfn-quicksight-dataset-datasetsemanticmetadata-custominstructions
	//
	CustomInstructions interface{} `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetsemanticmetadata.html#cfn-quicksight-dataset-datasetsemanticmetadata-description
	//
	Description interface{} `field:"optional" json:"description" yaml:"description"`
}

