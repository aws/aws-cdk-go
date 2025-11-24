package mixinsawsquicksight


// The range options for the data zoom scroll bar.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   visibleRangeOptionsProperty := &VisibleRangeOptionsProperty{
//   	PercentRange: &PercentVisibleRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visiblerangeoptions.html
//
type CfnTemplatePropsMixin_VisibleRangeOptionsProperty struct {
	// The percent range in the visible range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visiblerangeoptions.html#cfn-quicksight-template-visiblerangeoptions-percentrange
	//
	PercentRange interface{} `field:"optional" json:"percentRange" yaml:"percentRange"`
}

