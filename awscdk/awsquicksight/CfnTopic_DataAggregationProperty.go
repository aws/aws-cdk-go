package awsquicksight


// The definition of a data aggregation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataAggregationProperty := &DataAggregationProperty{
//   	DatasetRowDateGranularity: jsii.String("datasetRowDateGranularity"),
//   	DefaultDateColumnName: jsii.String("defaultDateColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-dataaggregation.html
//
type CfnTopic_DataAggregationProperty struct {
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-dataaggregation.html#cfn-quicksight-topic-dataaggregation-datasetrowdategranularity
	//
	DatasetRowDateGranularity *string `field:"optional" json:"datasetRowDateGranularity" yaml:"datasetRowDateGranularity"`
	// The column name for the default date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-dataaggregation.html#cfn-quicksight-topic-dataaggregation-defaultdatecolumnname
	//
	DefaultDateColumnName *string `field:"optional" json:"defaultDateColumnName" yaml:"defaultDateColumnName"`
}

