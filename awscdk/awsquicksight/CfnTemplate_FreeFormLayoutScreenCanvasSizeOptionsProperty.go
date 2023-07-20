package awsquicksight


// The options that determine the sizing of the canvas used in a free-form layout.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   freeFormLayoutScreenCanvasSizeOptionsProperty := &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   	OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutscreencanvassizeoptions.html
//
type CfnTemplate_FreeFormLayoutScreenCanvasSizeOptionsProperty struct {
	// The width that the view port will be optimized for when the layout renders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-freeformlayoutscreencanvassizeoptions.html#cfn-quicksight-template-freeformlayoutscreencanvassizeoptions-optimizedviewportwidth
	//
	OptimizedViewPortWidth *string `field:"required" json:"optimizedViewPortWidth" yaml:"optimizedViewPortWidth"`
}

