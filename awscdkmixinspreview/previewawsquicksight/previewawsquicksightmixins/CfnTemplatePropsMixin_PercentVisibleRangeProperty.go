package previewawsquicksightmixins


// The percent range in the visible range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   percentVisibleRangeProperty := &PercentVisibleRangeProperty{
//   	From: jsii.Number(123),
//   	To: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentvisiblerange.html
//
type CfnTemplatePropsMixin_PercentVisibleRangeProperty struct {
	// The lower bound of the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentvisiblerange.html#cfn-quicksight-template-percentvisiblerange-from
	//
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// The top bound of the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentvisiblerange.html#cfn-quicksight-template-percentvisiblerange-to
	//
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

