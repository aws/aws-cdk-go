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
type CfnTemplate_FreeFormLayoutScreenCanvasSizeOptionsProperty struct {
	// The width that the view port will be optimized for when the layout renders.
	OptimizedViewPortWidth *string `field:"required" json:"optimizedViewPortWidth" yaml:"optimizedViewPortWidth"`
}

