package awsquicksight


// A transform operation that changes the data types of one or more columns in the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   castColumnTypesOperationProperty := &CastColumnTypesOperationProperty{
//   	Alias: jsii.String("alias"),
//   	CastColumnTypeOperations: []interface{}{
//   		&CastColumnTypeOperationProperty{
//   			ColumnName: jsii.String("columnName"),
//   			NewColumnType: jsii.String("newColumnType"),
//
//   			// the properties below are optional
//   			Format: jsii.String("format"),
//   			SubType: jsii.String("subType"),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypesoperation.html
//
type CfnDataSet_CastColumnTypesOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypesoperation.html#cfn-quicksight-dataset-castcolumntypesoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The list of column type casting operations to perform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypesoperation.html#cfn-quicksight-dataset-castcolumntypesoperation-castcolumntypeoperations
	//
	CastColumnTypeOperations interface{} `field:"required" json:"castColumnTypeOperations" yaml:"castColumnTypeOperations"`
	// The source transform operation that provides input data for the type casting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-castcolumntypesoperation.html#cfn-quicksight-dataset-castcolumntypesoperation-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
}

