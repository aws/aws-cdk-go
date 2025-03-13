package awsquicksight


// The options that determine the sizing of the canvas used in a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutScreenCanvasSizeOptionsProperty := &GridLayoutScreenCanvasSizeOptionsProperty{
//   	ResizeOption: jsii.String("resizeOption"),
//
//   	// the properties below are optional
//   	OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutscreencanvassizeoptions.html
//
type CfnDashboard_GridLayoutScreenCanvasSizeOptionsProperty struct {
	// This value determines the layout behavior when the viewport is resized.
	//
	// - `FIXED` : A fixed width will be used when optimizing the layout. In the Amazon QuickSight console, this option is called `Classic` .
	// - `RESPONSIVE` : The width of the canvas will be responsive and optimized to the view port. In the Amazon QuickSight console, this option is called `Tiled` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutscreencanvassizeoptions.html#cfn-quicksight-dashboard-gridlayoutscreencanvassizeoptions-resizeoption
	//
	ResizeOption *string `field:"required" json:"resizeOption" yaml:"resizeOption"`
	// The width that the view port will be optimized for when the layout renders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutscreencanvassizeoptions.html#cfn-quicksight-dashboard-gridlayoutscreencanvassizeoptions-optimizedviewportwidth
	//
	OptimizedViewPortWidth *string `field:"optional" json:"optimizedViewPortWidth" yaml:"optimizedViewPortWidth"`
}

