package awsquicksight


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
type CfnDashboard_SectionBasedLayoutPaperCanvasSizeOptionsProperty struct {
	// `CfnDashboard.SectionBasedLayoutPaperCanvasSizeOptionsProperty.PaperMargin`.
	PaperMargin interface{} `field:"optional" json:"paperMargin" yaml:"paperMargin"`
	// `CfnDashboard.SectionBasedLayoutPaperCanvasSizeOptionsProperty.PaperOrientation`.
	PaperOrientation *string `field:"optional" json:"paperOrientation" yaml:"paperOrientation"`
	// `CfnDashboard.SectionBasedLayoutPaperCanvasSizeOptionsProperty.PaperSize`.
	PaperSize *string `field:"optional" json:"paperSize" yaml:"paperSize"`
}

