package previewawsquicksightmixins


// A transform operation that combines data from two sources based on specified join conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   joinOperationProperty := &JoinOperationProperty{
//   	Alias: jsii.String("alias"),
//   	LeftOperand: &TransformOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   	LeftOperandProperties: &JoinOperandPropertiesProperty{
//   		OutputColumnNameOverrides: []interface{}{
//   			&OutputColumnNameOverrideProperty{
//   				OutputColumnName: jsii.String("outputColumnName"),
//   				SourceColumnName: jsii.String("sourceColumnName"),
//   			},
//   		},
//   	},
//   	OnClause: jsii.String("onClause"),
//   	RightOperand: &TransformOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   	RightOperandProperties: &JoinOperandPropertiesProperty{
//   		OutputColumnNameOverrides: []interface{}{
//   			&OutputColumnNameOverrideProperty{
//   				OutputColumnName: jsii.String("outputColumnName"),
//   				SourceColumnName: jsii.String("sourceColumnName"),
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html
//
type CfnDataSetPropsMixin_JoinOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The left operand for the join operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-leftoperand
	//
	LeftOperand interface{} `field:"optional" json:"leftOperand" yaml:"leftOperand"`
	// Properties that control how the left operand's columns are handled in the join result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-leftoperandproperties
	//
	LeftOperandProperties interface{} `field:"optional" json:"leftOperandProperties" yaml:"leftOperandProperties"`
	// The join condition that specifies how to match rows between the left and right operands.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-onclause
	//
	OnClause *string `field:"optional" json:"onClause" yaml:"onClause"`
	// The right operand for the join operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-rightoperand
	//
	RightOperand interface{} `field:"optional" json:"rightOperand" yaml:"rightOperand"`
	// Properties that control how the right operand's columns are handled in the join result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-rightoperandproperties
	//
	RightOperandProperties interface{} `field:"optional" json:"rightOperandProperties" yaml:"rightOperandProperties"`
	// The type of join to perform, such as `INNER` , `LEFT` , `RIGHT` , or `OUTER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-joinoperation.html#cfn-quicksight-dataset-joinoperation-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

