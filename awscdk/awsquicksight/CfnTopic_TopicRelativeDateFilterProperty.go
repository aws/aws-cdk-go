package awsquicksight


// A structure that represents a relative date filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicRelativeDateFilterProperty := &TopicRelativeDateFilterProperty{
//   	Constant: &TopicSingularFilterConstantProperty{
//   		ConstantType: jsii.String("constantType"),
//   		SingularConstant: jsii.String("singularConstant"),
//   	},
//   	RelativeDateFilterFunction: jsii.String("relativeDateFilterFunction"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html
//
type CfnTopic_TopicRelativeDateFilterProperty struct {
	// The constant used in a relative date filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html#cfn-quicksight-topic-topicrelativedatefilter-constant
	//
	Constant interface{} `field:"optional" json:"constant" yaml:"constant"`
	// The function to be used in a relative date filter to determine the range of dates to include in the results.
	//
	// Valid values for this structure are `BEFORE` , `AFTER` , and `BETWEEN` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html#cfn-quicksight-topic-topicrelativedatefilter-relativedatefilterfunction
	//
	RelativeDateFilterFunction *string `field:"optional" json:"relativeDateFilterFunction" yaml:"relativeDateFilterFunction"`
	// The level of time precision that is used to aggregate `DateTime` values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html#cfn-quicksight-topic-topicrelativedatefilter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
}

