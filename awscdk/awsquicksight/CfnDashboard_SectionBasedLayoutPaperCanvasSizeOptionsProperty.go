package awsquicksight


// The options for a paper canvas of a section-based layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionBasedLayoutPaperCanvasSizeOptionsProperty := &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   	PaperMargin: &SpacingProperty{
//   		Bottom: jsii.String("bottom"),
//   		Left: jsii.String("left"),
//   		Right: jsii.String("right"),
//   		Top: jsii.String("top"),
//   	},
//   	PaperOrientation: jsii.String("paperOrientation"),
//   	PaperSize: jsii.String("paperSize"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions.html
//
type CfnDashboard_SectionBasedLayoutPaperCanvasSizeOptionsProperty struct {
	// Defines the spacing between the canvas content and the top, bottom, left, and right edges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions.html#cfn-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions-papermargin
	//
	PaperMargin interface{} `field:"optional" json:"paperMargin" yaml:"paperMargin"`
	// The paper orientation that is used to define canvas dimensions. Choose one of the following options:.
	//
	// - PORTRAIT
	// - LANDSCAPE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions.html#cfn-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions-paperorientation
	//
	PaperOrientation *string `field:"optional" json:"paperOrientation" yaml:"paperOrientation"`
	// The paper size that is used to define canvas dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions.html#cfn-quicksight-dashboard-sectionbasedlayoutpapercanvassizeoptions-papersize
	//
	PaperSize *string `field:"optional" json:"paperSize" yaml:"paperSize"`
}

