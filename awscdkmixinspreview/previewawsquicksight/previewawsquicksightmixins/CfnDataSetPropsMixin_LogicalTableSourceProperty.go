package previewawsquicksightmixins


// Information about the source of a logical table.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logicalTableSourceProperty := &LogicalTableSourceProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	JoinInstruction: &JoinInstructionProperty{
//   		LeftJoinKeyProperties: &JoinKeyPropertiesProperty{
//   			UniqueKey: jsii.Boolean(false),
//   		},
//   		LeftOperand: jsii.String("leftOperand"),
//   		OnClause: jsii.String("onClause"),
//   		RightJoinKeyProperties: &JoinKeyPropertiesProperty{
//   			UniqueKey: jsii.Boolean(false),
//   		},
//   		RightOperand: jsii.String("rightOperand"),
//   		Type: jsii.String("type"),
//   	},
//   	PhysicalTableId: jsii.String("physicalTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltablesource.html
//
type CfnDataSetPropsMixin_LogicalTableSourceProperty struct {
	// The Amazon Resource Number (ARN) of the parent dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltablesource.html#cfn-quicksight-dataset-logicaltablesource-datasetarn
	//
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// Specifies the result of a join of two logical tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltablesource.html#cfn-quicksight-dataset-logicaltablesource-joininstruction
	//
	JoinInstruction interface{} `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// Physical table ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-logicaltablesource.html#cfn-quicksight-dataset-logicaltablesource-physicaltableid
	//
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

