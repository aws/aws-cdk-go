package awsquicksight


// Configuration options for the canvas of a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutCanvasSizeOptionsProperty := &GridLayoutCanvasSizeOptionsProperty{
//   	ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   		ResizeOption: jsii.String("resizeOption"),
//
//   		// the properties below are optional
//   		OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutcanvassizeoptions.html
//
type CfnDashboard_GridLayoutCanvasSizeOptionsProperty struct {
	// The options that determine the sizing of the canvas used in a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gridlayoutcanvassizeoptions.html#cfn-quicksight-dashboard-gridlayoutcanvassizeoptions-screencanvassizeoptions
	//
	ScreenCanvasSizeOptions interface{} `field:"optional" json:"screenCanvasSizeOptions" yaml:"screenCanvasSizeOptions"`
}

