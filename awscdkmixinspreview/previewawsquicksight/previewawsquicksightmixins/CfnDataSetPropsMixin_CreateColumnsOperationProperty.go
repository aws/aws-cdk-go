package previewawsquicksightmixins


// A transform operation that creates calculated columns.
//
// Columns created in one such operation form a lexical closure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createColumnsOperationProperty := &CreateColumnsOperationProperty{
//   	Alias: jsii.String("alias"),
//   	Columns: []interface{}{
//   		&CalculatedColumnProperty{
//   			ColumnId: jsii.String("columnId"),
//   			ColumnName: jsii.String("columnName"),
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   	Source: &TransformOperationSourceProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-createcolumnsoperation.html
//
type CfnDataSetPropsMixin_CreateColumnsOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-createcolumnsoperation.html#cfn-quicksight-dataset-createcolumnsoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Calculated columns to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-createcolumnsoperation.html#cfn-quicksight-dataset-createcolumnsoperation-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The source transform operation that provides input data for creating new calculated columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-createcolumnsoperation.html#cfn-quicksight-dataset-createcolumnsoperation-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

