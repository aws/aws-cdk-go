package awsquicksight


// A transform operation that converts columns into rows, normalizing the data structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   unpivotOperationProperty := &UnpivotOperationProperty{
//   	Alias: jsii.String("alias"),
//   	ColumnsToUnpivot: []interface{}{
//   		&ColumnToUnpivotProperty{
//   			ColumnName: jsii.String("columnName"),
//   			NewValue: jsii.String("newValue"),
//   		},
//   	},
//   	Source: &TransformOperationSourceProperty{
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
//   	UnpivotedLabelColumnId: jsii.String("unpivotedLabelColumnId"),
//   	UnpivotedLabelColumnName: jsii.String("unpivotedLabelColumnName"),
//   	UnpivotedValueColumnId: jsii.String("unpivotedValueColumnId"),
//   	UnpivotedValueColumnName: jsii.String("unpivotedValueColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html
//
type CfnDataSet_UnpivotOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The list of columns to unpivot from the source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-columnstounpivot
	//
	ColumnsToUnpivot interface{} `field:"required" json:"columnsToUnpivot" yaml:"columnsToUnpivot"`
	// The source transform operation that provides input data for unpivoting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// A unique identifier for the new column that will contain the unpivoted column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-unpivotedlabelcolumnid
	//
	UnpivotedLabelColumnId *string `field:"required" json:"unpivotedLabelColumnId" yaml:"unpivotedLabelColumnId"`
	// The name for the new column that will contain the unpivoted column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-unpivotedlabelcolumnname
	//
	UnpivotedLabelColumnName *string `field:"required" json:"unpivotedLabelColumnName" yaml:"unpivotedLabelColumnName"`
	// A unique identifier for the new column that will contain the unpivoted values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-unpivotedvaluecolumnid
	//
	UnpivotedValueColumnId *string `field:"required" json:"unpivotedValueColumnId" yaml:"unpivotedValueColumnId"`
	// The name for the new column that will contain the unpivoted values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-unpivotoperation.html#cfn-quicksight-dataset-unpivotoperation-unpivotedvaluecolumnname
	//
	UnpivotedValueColumnName *string `field:"required" json:"unpivotedValueColumnName" yaml:"unpivotedValueColumnName"`
}

