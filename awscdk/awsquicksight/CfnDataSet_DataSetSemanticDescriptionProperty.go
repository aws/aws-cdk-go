package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetSemanticDescriptionProperty := &DataSetSemanticDescriptionProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetsemanticdescription.html
//
type CfnDataSet_DataSetSemanticDescriptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetsemanticdescription.html#cfn-quicksight-dataset-datasetsemanticdescription-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}

