package awsquicksight


// Specifies the source table and column mappings for an import table operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   importTableOperationSourceProperty := &ImportTableOperationSourceProperty{
//   	SourceTableId: jsii.String("sourceTableId"),
//
//   	// the properties below are optional
//   	ColumnIdMappings: []interface{}{
//   		&DataSetColumnIdMappingProperty{
//   			SourceColumnId: jsii.String("sourceColumnId"),
//   			TargetColumnId: jsii.String("targetColumnId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-importtableoperationsource.html
//
type CfnDataSet_ImportTableOperationSourceProperty struct {
	// The identifier of the source table to import data from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-importtableoperationsource.html#cfn-quicksight-dataset-importtableoperationsource-sourcetableid
	//
	SourceTableId *string `field:"required" json:"sourceTableId" yaml:"sourceTableId"`
	// The mappings between source column identifiers and target column identifiers during the import.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-importtableoperationsource.html#cfn-quicksight-dataset-importtableoperationsource-columnidmappings
	//
	ColumnIdMappings interface{} `field:"optional" json:"columnIdMappings" yaml:"columnIdMappings"`
}

