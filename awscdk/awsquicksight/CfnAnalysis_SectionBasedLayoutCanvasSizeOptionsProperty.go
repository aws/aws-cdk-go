package awsquicksight


// The options for the canvas of a section-based layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sectionBasedLayoutCanvasSizeOptionsProperty := &SectionBasedLayoutCanvasSizeOptionsProperty{
//   	PaperCanvasSizeOptions: &SectionBasedLayoutPaperCanvasSizeOptionsProperty{
//   		PaperMargin: &SpacingProperty{
//   			Bottom: jsii.String("bottom"),
//   			Left: jsii.String("left"),
//   			Right: jsii.String("right"),
//   			Top: jsii.String("top"),
//   		},
//   		PaperOrientation: jsii.String("paperOrientation"),
//   		PaperSize: jsii.String("paperSize"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutcanvassizeoptions.html
//
type CfnAnalysis_SectionBasedLayoutCanvasSizeOptionsProperty struct {
	// The options for a paper canvas of a section-based layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutcanvassizeoptions.html#cfn-quicksight-analysis-sectionbasedlayoutcanvassizeoptions-papercanvassizeoptions
	//
	PaperCanvasSizeOptions interface{} `field:"optional" json:"paperCanvasSizeOptions" yaml:"paperCanvasSizeOptions"`
}

