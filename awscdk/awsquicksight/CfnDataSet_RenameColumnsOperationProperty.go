package awsquicksight


// A transform operation that renames one or more columns in the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameColumnsOperationProperty := &RenameColumnsOperationProperty{
//   	Alias: jsii.String("alias"),
//   	RenameColumnOperations: []interface{}{
//   		&RenameColumnOperationProperty{
//   			ColumnName: jsii.String("columnName"),
//   			NewColumnName: jsii.String("newColumnName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnsoperation.html
//
type CfnDataSet_RenameColumnsOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnsoperation.html#cfn-quicksight-dataset-renamecolumnsoperation-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The list of column rename operations to perform, specifying old and new column names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnsoperation.html#cfn-quicksight-dataset-renamecolumnsoperation-renamecolumnoperations
	//
	RenameColumnOperations interface{} `field:"required" json:"renameColumnOperations" yaml:"renameColumnOperations"`
	// The source transform operation that provides input data for column renaming.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-renamecolumnsoperation.html#cfn-quicksight-dataset-renamecolumnsoperation-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
}

