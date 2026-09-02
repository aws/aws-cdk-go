package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetRelationEndpointProperty := &DataSetRelationEndpointProperty{
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	DataSetArn: jsii.String("dataSetArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelationendpoint.html
//
type CfnTopicV2_DataSetRelationEndpointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelationendpoint.html#cfn-quicksight-topicv2-datasetrelationendpoint-columnnames
	//
	ColumnNames *[]*string `field:"required" json:"columnNames" yaml:"columnNames"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelationendpoint.html#cfn-quicksight-topicv2-datasetrelationendpoint-datasetarn
	//
	DataSetArn *string `field:"required" json:"dataSetArn" yaml:"dataSetArn"`
}

