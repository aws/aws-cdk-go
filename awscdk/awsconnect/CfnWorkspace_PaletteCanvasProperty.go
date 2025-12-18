package awsconnect


// Contains color configuration for canvas elements in a workspace theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   paletteCanvasProperty := &PaletteCanvasProperty{
//   	ActiveBackground: jsii.String("activeBackground"),
//   	ContainerBackground: jsii.String("containerBackground"),
//   	PageBackground: jsii.String("pageBackground"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettecanvas.html
//
type CfnWorkspace_PaletteCanvasProperty struct {
	// The background color for active elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettecanvas.html#cfn-connect-workspace-palettecanvas-activebackground
	//
	ActiveBackground *string `field:"optional" json:"activeBackground" yaml:"activeBackground"`
	// The background color for container elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettecanvas.html#cfn-connect-workspace-palettecanvas-containerbackground
	//
	ContainerBackground *string `field:"optional" json:"containerBackground" yaml:"containerBackground"`
	// The background color for page elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettecanvas.html#cfn-connect-workspace-palettecanvas-pagebackground
	//
	PageBackground *string `field:"optional" json:"pageBackground" yaml:"pageBackground"`
}

