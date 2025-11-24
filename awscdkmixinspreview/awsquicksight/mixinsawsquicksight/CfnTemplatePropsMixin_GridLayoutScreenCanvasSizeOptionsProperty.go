package mixinsawsquicksight


// The options that determine the sizing of the canvas used in a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gridLayoutScreenCanvasSizeOptionsProperty := &GridLayoutScreenCanvasSizeOptionsProperty{
//   	OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   	ResizeOption: jsii.String("resizeOption"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutscreencanvassizeoptions.html
//
type CfnTemplatePropsMixin_GridLayoutScreenCanvasSizeOptionsProperty struct {
	// The width that the view port will be optimized for when the layout renders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutscreencanvassizeoptions.html#cfn-quicksight-template-gridlayoutscreencanvassizeoptions-optimizedviewportwidth
	//
	OptimizedViewPortWidth *string `field:"optional" json:"optimizedViewPortWidth" yaml:"optimizedViewPortWidth"`
	// This value determines the layout behavior when the viewport is resized.
	//
	// - `FIXED` : A fixed width will be used when optimizing the layout. In the Quick Sight console, this option is called `Classic` .
	// - `RESPONSIVE` : The width of the canvas will be responsive and optimized to the view port. In the Quick Sight console, this option is called `Tiled` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutscreencanvassizeoptions.html#cfn-quicksight-template-gridlayoutscreencanvassizeoptions-resizeoption
	//
	ResizeOption *string `field:"optional" json:"resizeOption" yaml:"resizeOption"`
}

