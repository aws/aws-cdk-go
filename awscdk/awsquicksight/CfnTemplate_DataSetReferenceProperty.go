package awsquicksight


// Dataset reference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetReferenceProperty := &DataSetReferenceProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	DataSetPlaceholder: jsii.String("dataSetPlaceholder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetreference.html
//
type CfnTemplate_DataSetReferenceProperty struct {
	// Dataset Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetreference.html#cfn-quicksight-template-datasetreference-datasetarn
	//
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
	// Dataset placeholder.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetreference.html#cfn-quicksight-template-datasetreference-datasetplaceholder
	//
	DataSetPlaceholder *string `field:"required" json:"dataSetPlaceholder" yaml:"dataSetPlaceholder"`
}

