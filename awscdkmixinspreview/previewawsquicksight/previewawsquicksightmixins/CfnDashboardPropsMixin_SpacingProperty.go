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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-spacing.html
//
type CfnDashboardPropsMixin_SpacingProperty struct {
	// Define the bottom spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-spacing.html#cfn-quicksight-dashboard-spacing-bottom
	//
	Bottom *string `field:"optional" json:"bottom" yaml:"bottom"`
	// Define the left spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-spacing.html#cfn-quicksight-dashboard-spacing-left
	//
	Left *string `field:"optional" json:"left" yaml:"left"`
	// Define the right spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-spacing.html#cfn-quicksight-dashboard-spacing-right
	//
	Right *string `field:"optional" json:"right" yaml:"right"`
	// Define the top spacing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-spacing.html#cfn-quicksight-dashboard-spacing-top
	//
	Top *string `field:"optional" json:"top" yaml:"top"`
}

