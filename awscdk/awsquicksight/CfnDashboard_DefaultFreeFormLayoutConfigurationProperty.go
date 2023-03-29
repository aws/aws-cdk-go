package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultFreeFormLayoutConfigurationProperty := &DefaultFreeFormLayoutConfigurationProperty{
//   	CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   		ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   			OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   		},
//   	},
//   }
//
type CfnDashboard_DefaultFreeFormLayoutConfigurationProperty struct {
	// `CfnDashboard.DefaultFreeFormLayoutConfigurationProperty.CanvasSizeOptions`.
	CanvasSizeOptions interface{} `field:"required" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

