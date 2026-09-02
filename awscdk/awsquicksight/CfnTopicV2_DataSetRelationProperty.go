package awsquicksight


// A relation between two datasets referenced by a V2 (SEMANTIC_VIEW) topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSetRelationProperty := &DataSetRelationProperty{
//   	Left: &DataSetRelationEndpointProperty{
//   		ColumnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		DataSetArn: jsii.String("dataSetArn"),
//   	},
//   	Right: &DataSetRelationEndpointProperty{
//   		ColumnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		DataSetArn: jsii.String("dataSetArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelation.html
//
type CfnTopicV2_DataSetRelationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelation.html#cfn-quicksight-topicv2-datasetrelation-left
	//
	Left interface{} `field:"required" json:"left" yaml:"left"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topicv2-datasetrelation.html#cfn-quicksight-topicv2-datasetrelation-right
	//
	Right interface{} `field:"required" json:"right" yaml:"right"`
}

