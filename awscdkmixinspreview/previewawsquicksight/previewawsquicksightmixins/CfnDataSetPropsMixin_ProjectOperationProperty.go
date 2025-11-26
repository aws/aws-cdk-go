package previewawsquicksightmixins


// A transform operation that projects columns.
//
// Operations that come after a projection can only refer to projected columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   projectOperationProperty := &ProjectOperationProperty{
//   	Alias: jsii.String("alias"),
//   	ProjectedColumns: []*string{
//   		jsii.String("projectedColumns"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-projectoperation.html
//
type CfnDataSetPropsMixin_ProjectOperationProperty struct {
	// Alias for this operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-projectoperation.html#cfn-quicksight-dataset-projectoperation-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Projected columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-projectoperation.html#cfn-quicksight-dataset-projectoperation-projectedcolumns
	//
	ProjectedColumns *[]*string `field:"optional" json:"projectedColumns" yaml:"projectedColumns"`
	// The source transform operation that provides input data for column projection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-projectoperation.html#cfn-quicksight-dataset-projectoperation-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}

