package awsquicksight


// A filter that filters topics based on the value of a numeric field.
//
// The filter includes only topics whose numeric field value falls within the specified range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicNumericRangeFilterProperty := &TopicNumericRangeFilterProperty{
//   	Aggregation: jsii.String("aggregation"),
//   	Constant: &TopicRangeFilterConstantProperty{
//   		ConstantType: jsii.String("constantType"),
//   		RangeConstant: &RangeConstantProperty{
//   			Maximum: jsii.String("maximum"),
//   			Minimum: jsii.String("minimum"),
//   		},
//   	},
//   	Inclusive: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnumericrangefilter.html
//
type CfnTopic_TopicNumericRangeFilterProperty struct {
	// An aggregation function that specifies how to calculate the value of a numeric field for a topic, Valid values for this structure are `NO_AGGREGATION` , `SUM` , `AVERAGE` , `COUNT` , `DISTINCT_COUNT` , `MAX` , `MEDIAN` , `MIN` , `STDEV` , `STDEVP` , `VAR` , and `VARP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnumericrangefilter.html#cfn-quicksight-topic-topicnumericrangefilter-aggregation
	//
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The constant used in a numeric range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnumericrangefilter.html#cfn-quicksight-topic-topicnumericrangefilter-constant
	//
	Constant interface{} `field:"optional" json:"constant" yaml:"constant"`
	// A Boolean value that indicates whether the endpoints of the numeric range are included in the filter.
	//
	// If set to true, topics whose numeric field value is equal to the endpoint values will be included in the filter. If set to false, topics whose numeric field value is equal to the endpoint values will be excluded from the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicnumericrangefilter.html#cfn-quicksight-topic-topicnumericrangefilter-inclusive
	//
	// Default: - false.
	//
	Inclusive interface{} `field:"optional" json:"inclusive" yaml:"inclusive"`
}

