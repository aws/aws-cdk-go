package previewawsquicksightmixins


// A transform operation that tags a column with additional information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagColumnOperationProperty := &TagColumnOperationProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Tags: []ColumnTagProperty{
//   		&ColumnTagProperty{
//   			ColumnDescription: &ColumnDescriptionProperty{
//   				Text: jsii.String("text"),
//   			},
//   			ColumnGeographicRole: jsii.String("columnGeographicRole"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tagcolumnoperation.html
//
type CfnDataSetPropsMixin_TagColumnOperationProperty struct {
	// The column that this operation acts on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tagcolumnoperation.html#cfn-quicksight-dataset-tagcolumnoperation-columnname
	//
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// The dataset column tag, currently only used for geospatial type tagging.
	//
	// > This is not tags for the AWS tagging feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-tagcolumnoperation.html#cfn-quicksight-dataset-tagcolumnoperation-tags
	//
	Tags *[]*CfnDataSetPropsMixin_ColumnTagProperty `field:"optional" json:"tags" yaml:"tags"`
}

