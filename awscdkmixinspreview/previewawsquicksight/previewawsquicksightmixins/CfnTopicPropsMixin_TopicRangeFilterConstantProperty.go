package previewawsquicksightmixins


// A constant value that is used in a range filter to specify the endpoints of the range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicRangeFilterConstantProperty := &TopicRangeFilterConstantProperty{
//   	ConstantType: jsii.String("constantType"),
//   	RangeConstant: &RangeConstantProperty{
//   		Maximum: jsii.String("maximum"),
//   		Minimum: jsii.String("minimum"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrangefilterconstant.html
//
type CfnTopicPropsMixin_TopicRangeFilterConstantProperty struct {
	// The data type of the constant value that is used in a range filter.
	//
	// Valid values for this structure are `RANGE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrangefilterconstant.html#cfn-quicksight-topic-topicrangefilterconstant-constanttype
	//
	ConstantType *string `field:"optional" json:"constantType" yaml:"constantType"`
	// The value of the constant that is used to specify the endpoints of a range filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrangefilterconstant.html#cfn-quicksight-topic-topicrangefilterconstant-rangeconstant
	//
	RangeConstant interface{} `field:"optional" json:"rangeConstant" yaml:"rangeConstant"`
}

