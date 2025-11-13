package awsquicksight


// A transform operation that combines rows from two data sources by stacking them vertically (union operation).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   appendOperationProperty := &AppendOperationProperty{
//   	Alias: jsii.String("alias"),
//   	AppendedColumns: []interface{}{
//   		&AppendedColumnProperty{
//   			ColumnName: jsii.String("columnName"),
//   			NewColumnId: jsii.String("newColumnId"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	FirstSource: &TransformOperationSourceProperty{
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
//   	SecondSource: &TransformOperationSourceProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html
//
type CfnDataSet_AppendOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The list of columns to include in the appended result, mapping columns from both sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-appendedcolumns
	//
	AppendedColumns interface{} `field:"required" json:"appendedColumns" yaml:"appendedColumns"`
	// The first data source to be included in the append operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-firstsource
	//
	FirstSource interface{} `field:"optional" json:"firstSource" yaml:"firstSource"`
	// The second data source to be appended to the first source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-secondsource
	//
	SecondSource interface{} `field:"optional" json:"secondSource" yaml:"secondSource"`
}

