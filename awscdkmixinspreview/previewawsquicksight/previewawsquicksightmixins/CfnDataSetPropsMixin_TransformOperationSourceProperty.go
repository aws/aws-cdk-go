package previewawsquicksightmixins


// Specifies the source of data for a transform operation, including the source operation and column mappings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformOperationSourceProperty := &TransformOperationSourceProperty{
//   	ColumnIdMappings: []interface{}{
//   		&DataSetColumnIdMappingProperty{
//   			SourceColumnId: jsii.String("sourceColumnId"),
//   			TargetColumnId: jsii.String("targetColumnId"),
//   		},
//   	},
//   	TransformOperationId: jsii.String("transformOperationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperationsource.html
//
type CfnDataSetPropsMixin_TransformOperationSourceProperty struct {
	// The mappings between source column identifiers and target column identifiers for this transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperationsource.html#cfn-quicksight-dataset-transformoperationsource-columnidmappings
	//
	ColumnIdMappings interface{} `field:"optional" json:"columnIdMappings" yaml:"columnIdMappings"`
	// The identifier of the transform operation that provides input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-transformoperationsource.html#cfn-quicksight-dataset-transformoperationsource-transformoperationid
	//
	TransformOperationId *string `field:"optional" json:"transformOperationId" yaml:"transformOperationId"`
}

