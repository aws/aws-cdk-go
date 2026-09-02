package awsquicksight


// A dataset reference used by a V2 (SEMANTIC_VIEW) topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataSetReferenceProperty := &DataSetReferenceProperty{
//   	DataSetArn: jsii.String("dataSetArn"),
//   	DataSetName: jsii.String("dataSetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetreference.html
//
type CfnTopicV2PropsMixin_DataSetReferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetreference.html#cfn-quicksight-topicv2-datasetreference-datasetarn
	//
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetreference.html#cfn-quicksight-topicv2-datasetreference-datasetname
	//
	DataSetName *string `field:"optional" json:"dataSetName" yaml:"dataSetName"`
}

