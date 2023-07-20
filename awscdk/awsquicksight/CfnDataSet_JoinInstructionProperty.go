package awsquicksight


// The instructions associated with a join.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinInstructionProperty := &JoinInstructionProperty{
//   	LeftOperand: jsii.String("leftOperand"),
//   	OnClause: jsii.String("onClause"),
//   	RightOperand: jsii.String("rightOperand"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	LeftJoinKeyProperties: &JoinKeyPropertiesProperty{
//   		UniqueKey: jsii.Boolean(false),
//   	},
//   	RightJoinKeyProperties: &JoinKeyPropertiesProperty{
//   		UniqueKey: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html
//
type CfnDataSet_JoinInstructionProperty struct {
	// The operand on the left side of a join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-leftoperand
	//
	LeftOperand *string `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// The join instructions provided in the `ON` clause of a join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-onclause
	//
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// The operand on the right side of a join.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-rightoperand
	//
	RightOperand *string `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// The type of join that it is.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Join key properties of the left operand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-leftjoinkeyproperties
	//
	LeftJoinKeyProperties interface{} `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// Join key properties of the right operand.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joininstruction.html#cfn-quicksight-dataset-joininstruction-rightjoinkeyproperties
	//
	RightJoinKeyProperties interface{} `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
}

