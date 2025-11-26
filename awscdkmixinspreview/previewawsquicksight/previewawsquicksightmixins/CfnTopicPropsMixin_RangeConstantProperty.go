package previewawsquicksightmixins


// The value of the constant that is used to specify the endpoints of a range filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rangeConstantProperty := &RangeConstantProperty{
//   	Maximum: jsii.String("maximum"),
//   	Minimum: jsii.String("minimum"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-rangeconstant.html
//
type CfnTopicPropsMixin_RangeConstantProperty struct {
	// The maximum value for a range constant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-rangeconstant.html#cfn-quicksight-topic-rangeconstant-maximum
	//
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// The minimum value for a range constant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-rangeconstant.html#cfn-quicksight-topic-rangeconstant-minimum
	//
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
}

