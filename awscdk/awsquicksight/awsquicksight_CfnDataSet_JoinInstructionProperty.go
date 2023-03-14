package awsquicksight


// The instructions associated with a join.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinInstructionProperty := &joinInstructionProperty{
//   	leftOperand: jsii.String("leftOperand"),
//   	onClause: jsii.String("onClause"),
//   	rightOperand: jsii.String("rightOperand"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	leftJoinKeyProperties: &joinKeyPropertiesProperty{
//   		uniqueKey: jsii.Boolean(false),
//   	},
//   	rightJoinKeyProperties: &joinKeyPropertiesProperty{
//   		uniqueKey: jsii.Boolean(false),
//   	},
//   }
//
type CfnDataSet_JoinInstructionProperty struct {
	// The operand on the left side of a join.
	LeftOperand *string `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// The join instructions provided in the `ON` clause of a join.
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// The operand on the right side of a join.
	RightOperand *string `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// The type of join that it is.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Join key properties of the left operand.
	LeftJoinKeyProperties interface{} `field:"optional" json:"leftJoinKeyProperties" yaml:"leftJoinKeyProperties"`
	// Join key properties of the right operand.
	RightJoinKeyProperties interface{} `field:"optional" json:"rightJoinKeyProperties" yaml:"rightJoinKeyProperties"`
}

