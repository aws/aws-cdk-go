package mixinsawsquicksight


// A transform operation that combines rows from two data sources by stacking them vertically (union operation).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   appendOperationProperty := &AppendOperationProperty{
//   	Alias: jsii.String("alias"),
//   	AppendedColumns: []interface{}{
//   		&AppendedColumnProperty{
//   			ColumnName: jsii.String("columnName"),
//   			NewColumnId: jsii.String("newColumnId"),
//   		},
//   	},
//   	FirstSource: &TransformOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   	SecondSource: &TransformOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		TransformOperationId: jsii.String("transformOperationId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html
//
type CfnDataSetPropsMixin_AppendOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The list of columns to include in the appended result, mapping columns from both sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-appendedcolumns
	//
	AppendedColumns interface{} `field:"optional" json:"appendedColumns" yaml:"appendedColumns"`
	// The first data source to be included in the append operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-firstsource
	//
	FirstSource interface{} `field:"optional" json:"firstSource" yaml:"firstSource"`
	// The second data source to be appended to the first source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-appendoperation.html#cfn-quicksight-dataset-appendoperation-secondsource
	//
	SecondSource interface{} `field:"optional" json:"secondSource" yaml:"secondSource"`
}

