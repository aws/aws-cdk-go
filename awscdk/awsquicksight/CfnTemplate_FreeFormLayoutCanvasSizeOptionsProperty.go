package awsquicksight


// Configuration options for the canvas of a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeFormLayoutCanvasSizeOptionsProperty := &FreeFormLayoutCanvasSizeOptionsProperty{
//   	ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   		OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutcanvassizeoptions.html
//
type CfnTemplate_FreeFormLayoutCanvasSizeOptionsProperty struct {
	// The options that determine the sizing of the canvas used in a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutcanvassizeoptions.html#cfn-quicksight-template-freeformlayoutcanvassizeoptions-screencanvassizeoptions
	//
	ScreenCanvasSizeOptions interface{} `field:"optional" json:"screenCanvasSizeOptions" yaml:"screenCanvasSizeOptions"`
}

