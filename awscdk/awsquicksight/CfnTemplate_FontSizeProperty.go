package awsquicksight


// The option that determines the text display size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fontSizeProperty := &FontSizeProperty{
//   	Absolute: jsii.String("absolute"),
//   	Relative: jsii.String("relative"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontsize.html
//
type CfnTemplate_FontSizeProperty struct {
	// The font size that you want to use in px.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontsize.html#cfn-quicksight-template-fontsize-absolute
	//
	Absolute *string `field:"optional" json:"absolute" yaml:"absolute"`
	// The lexical name for the text size, proportional to its surrounding context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontsize.html#cfn-quicksight-template-fontsize-relative
	//
	Relative *string `field:"optional" json:"relative" yaml:"relative"`
}

