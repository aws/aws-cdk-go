package awsquicksight


// A filter used to restrict data based on a range of dates or times.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicDateRangeFilterProperty := &TopicDateRangeFilterProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicdaterangefilter.html
//
type CfnTopic_TopicDateRangeFilterProperty struct {
	// The constant used in a date range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicdaterangefilter.html#cfn-quicksight-topic-topicdaterangefilter-constant
	//
	Constant interface{} `field:"optional" json:"constant" yaml:"constant"`
	// A Boolean value that indicates whether the date range filter should include the boundary values.
	//
	// If set to true, the filter includes the start and end dates. If set to false, the filter excludes them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicdaterangefilter.html#cfn-quicksight-topic-topicdaterangefilter-inclusive
	//
	Inclusive interface{} `field:"optional" json:"inclusive" yaml:"inclusive"`
}

