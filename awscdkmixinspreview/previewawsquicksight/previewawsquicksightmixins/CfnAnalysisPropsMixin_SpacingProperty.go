package previewawsquicksightmixins


// The configuration of spacing (often a margin or padding).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   spacingProperty := &SpacingProperty{
//   	Bottom: jsii.String("bottom"),
//   	Left: jsii.String("left"),
//   	Right: jsii.String("right"),
//   	Top: jsii.String("top"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spacing.html
//
type CfnAnalysisPropsMixin_SpacingProperty struct {
	// Define the bottom spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spacing.html#cfn-quicksight-analysis-spacing-bottom
	//
	Bottom *string `field:"optional" json:"bottom" yaml:"bottom"`
	// Define the left spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spacing.html#cfn-quicksight-analysis-spacing-left
	//
	Left *string `field:"optional" json:"left" yaml:"left"`
	// Define the right spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spacing.html#cfn-quicksight-analysis-spacing-right
	//
	Right *string `field:"optional" json:"right" yaml:"right"`
	// Define the top spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-spacing.html#cfn-quicksight-analysis-spacing-top
	//
	Top *string `field:"optional" json:"top" yaml:"top"`
}

