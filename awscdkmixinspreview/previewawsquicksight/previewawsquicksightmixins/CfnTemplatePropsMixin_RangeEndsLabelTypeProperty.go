package previewawsquicksightmixins


// The range ends label type of a data path label.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rangeEndsLabelTypeProperty := &RangeEndsLabelTypeProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rangeendslabeltype.html
//
type CfnTemplatePropsMixin_RangeEndsLabelTypeProperty struct {
	// The visibility of the range ends label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rangeendslabeltype.html#cfn-quicksight-template-rangeendslabeltype-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

