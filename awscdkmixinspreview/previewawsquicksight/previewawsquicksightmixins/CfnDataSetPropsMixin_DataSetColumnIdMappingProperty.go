package previewawsquicksightmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetColumnIdMappingProperty := &DataSetColumnIdMappingProperty{
//   	SourceColumnId: jsii.String("sourceColumnId"),
//   	TargetColumnId: jsii.String("targetColumnId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetcolumnidmapping.html
//
type CfnDataSetPropsMixin_DataSetColumnIdMappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetcolumnidmapping.html#cfn-quicksight-dataset-datasetcolumnidmapping-sourcecolumnid
	//
	SourceColumnId *string `field:"optional" json:"sourceColumnId" yaml:"sourceColumnId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetcolumnidmapping.html#cfn-quicksight-dataset-datasetcolumnidmapping-targetcolumnid
	//
	TargetColumnId *string `field:"optional" json:"targetColumnId" yaml:"targetColumnId"`
}

