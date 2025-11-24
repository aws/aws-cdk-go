package mixinsawsquicksight


// Configuration options for the canvas of a grid layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gridLayoutCanvasSizeOptionsProperty := &GridLayoutCanvasSizeOptionsProperty{
//   	ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   		OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   		ResizeOption: jsii.String("resizeOption"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutcanvassizeoptions.html
//
type CfnAnalysisPropsMixin_GridLayoutCanvasSizeOptionsProperty struct {
	// The options that determine the sizing of the canvas used in a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutcanvassizeoptions.html#cfn-quicksight-analysis-gridlayoutcanvassizeoptions-screencanvassizeoptions
	//
	ScreenCanvasSizeOptions interface{} `field:"optional" json:"screenCanvasSizeOptions" yaml:"screenCanvasSizeOptions"`
}

