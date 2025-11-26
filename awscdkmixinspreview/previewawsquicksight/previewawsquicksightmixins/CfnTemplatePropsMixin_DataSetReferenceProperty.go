package previewawsquicksightmixins


// Dataset reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetReferenceProperty := &DataSetReferenceProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetreference.html
//
type CfnTemplatePropsMixin_DataSetReferenceProperty struct {
	// Dataset Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetreference.html#cfn-quicksight-template-datasetreference-datasetarn
	//
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// Dataset placeholder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetreference.html#cfn-quicksight-template-datasetreference-datasetplaceholder
	//
	DataSetPlaceholder *string `field:"optional" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

