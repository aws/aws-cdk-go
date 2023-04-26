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
type CfnTemplate_FreeFormLayoutCanvasSizeOptionsProperty struct {
	// The options that determine the sizing of the canvas used in a free-form layout.
	ScreenCanvasSizeOptions interface{} `field:"optional" json:"screenCanvasSizeOptions" yaml:"screenCanvasSizeOptions"`
}

