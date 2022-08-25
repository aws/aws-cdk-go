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
//   logicalTableSourceProperty := &logicalTableSourceProperty{
//   	dataSetArn: jsii.String("dataSetArn"),
//   	joinInstruction: &joinInstructionProperty{
//   		leftOperand: jsii.String("leftOperand"),
//   		onClause: jsii.String("onClause"),
//   		rightOperand: jsii.String("rightOperand"),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		leftJoinKeyProperties: &joinKeyPropertiesProperty{
//   			uniqueKey: jsii.Boolean(false),
//   		},
//   		rightJoinKeyProperties: &joinKeyPropertiesProperty{
//   			uniqueKey: jsii.Boolean(false),
//   		},
//   	},
//   	physicalTableId: jsii.String("physicalTableId"),
//   }
//
type CfnDataSet_LogicalTableSourceProperty struct {
	// `CfnDataSet.LogicalTableSourceProperty.DataSetArn`.
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// Specifies the result of a join of two logical tables.
	JoinInstruction interface{} `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// Physical table ID.
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

