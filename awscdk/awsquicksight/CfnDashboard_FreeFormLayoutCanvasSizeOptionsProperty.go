package awsquicksight


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
type CfnDashboard_FreeFormLayoutCanvasSizeOptionsProperty struct {
	// `CfnDashboard.FreeFormLayoutCanvasSizeOptionsProperty.ScreenCanvasSizeOptions`.
	ScreenCanvasSizeOptions interface{} `field:"optional" json:"screenCanvasSizeOptions" yaml:"screenCanvasSizeOptions"`
}

