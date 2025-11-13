package awsquicksight


// A transform operation that combines data from two sources based on specified join conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   joinOperationProperty := &JoinOperationProperty{
//   	Alias: jsii.String("alias"),
//   	LeftOperand: &TransformOperationSourceProperty{
//   		TransformOperationId: jsii.String("transformOperationId"),
//
//   		// the properties below are optional
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   	},
//   	OnClause: jsii.String("onClause"),
//   	RightOperand: &TransformOperationSourceProperty{
//   		TransformOperationId: jsii.String("transformOperationId"),
//
//   		// the properties below are optional
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	LeftOperandProperties: &JoinOperandPropertiesProperty{
//   		OutputColumnNameOverrides: []interface{}{
//   			&OutputColumnNameOverrideProperty{
//   				OutputColumnName: jsii.String("outputColumnName"),
//
//   				// the properties below are optional
//   				SourceColumnName: jsii.String("sourceColumnName"),
//   			},
//   		},
//   	},
//   	RightOperandProperties: &JoinOperandPropertiesProperty{
//   		OutputColumnNameOverrides: []interface{}{
//   			&OutputColumnNameOverrideProperty{
//   				OutputColumnName: jsii.String("outputColumnName"),
//
//   				// the properties below are optional
//   				SourceColumnName: jsii.String("sourceColumnName"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html
//
type CfnDataSet_JoinOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The left operand for the join operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-leftoperand
	//
	LeftOperand interface{} `field:"required" json:"leftOperand" yaml:"leftOperand"`
	// The join condition that specifies how to match rows between the left and right operands.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-onclause
	//
	OnClause *string `field:"required" json:"onClause" yaml:"onClause"`
	// The right operand for the join operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-rightoperand
	//
	RightOperand interface{} `field:"required" json:"rightOperand" yaml:"rightOperand"`
	// The type of join to perform, such as `INNER` , `LEFT` , `RIGHT` , or `OUTER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Properties that control how the left operand's columns are handled in the join result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-leftoperandproperties
	//
	LeftOperandProperties interface{} `field:"optional" json:"leftOperandProperties" yaml:"leftOperandProperties"`
	// Properties that control how the right operand's columns are handled in the join result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-rightoperandproperties
	//
	RightOperandProperties interface{} `field:"optional" json:"rightOperandProperties" yaml:"rightOperandProperties"`
}

