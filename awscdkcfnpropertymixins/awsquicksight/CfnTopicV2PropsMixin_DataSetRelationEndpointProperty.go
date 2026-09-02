package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnTopicV2PropsMixin_DataSetRelationEndpointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelationendpoint.html#cfn-quicksight-topicv2-datasetrelationendpoint-columnnames
	//
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelationendpoint.html#cfn-quicksight-topicv2-datasetrelationendpoint-datasetarn
	//
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
}

