package awsquicksight


// Information about the source of a logical table.
//
// This is a variant type structure. For this structure to be valid, only one of the attributes can be non-null.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logicalTableSourceProperty := &LogicalTableSourceProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	JoinInstruction: &JoinInstructionProperty{
//   		LeftOperand: jsii.String("leftOperand"),
//   		OnClause: jsii.String("onClause"),
//   		RightOperand: jsii.String("rightOperand"),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		LeftJoinKeyProperties: &JoinKeyPropertiesProperty{
//   			UniqueKey: jsii.Boolean(false),
//   		},
//   		RightJoinKeyProperties: &JoinKeyPropertiesProperty{
//   			UniqueKey: jsii.Boolean(false),
//   		},
//   	},
//   	PhysicalTableId: jsii.String("physicalTableId"),
//   }
//
type CfnDataSet_LogicalTableSourceProperty struct {
	// The Amazon Resource Number (ARN) of the parent dataset.
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// Specifies the result of a join of two logical tables.
	JoinInstruction interface{} `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// Physical table ID.
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

