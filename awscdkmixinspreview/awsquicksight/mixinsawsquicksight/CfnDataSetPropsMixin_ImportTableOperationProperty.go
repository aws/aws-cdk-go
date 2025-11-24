package mixinsawsquicksight


// A transform operation that imports data from a source table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importTableOperationProperty := &ImportTableOperationProperty{
//   	Alias: jsii.String("alias"),
//   	Source: &ImportTableOperationSourceProperty{
//   		ColumnIdMappings: []interface{}{
//   			&DataSetColumnIdMappingProperty{
//   				SourceColumnId: jsii.String("sourceColumnId"),
//   				TargetColumnId: jsii.String("targetColumnId"),
//   			},
//   		},
//   		SourceTableId: jsii.String("sourceTableId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-importtableoperation.html
//
type CfnDataSetPropsMixin_ImportTableOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-importtableoperation.html#cfn-quicksight-dataset-importtableoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The source configuration that specifies which source table to import and any column mappings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-importtableoperation.html#cfn-quicksight-dataset-importtableoperation-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

