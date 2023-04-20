package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultGridLayoutConfigurationProperty := &DefaultGridLayoutConfigurationProperty{
//   	CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   		ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   			ResizeOption: jsii.String("resizeOption"),
//
//   			// the properties below are optional
//   			OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   		},
//   	},
//   }
//
type CfnTemplate_DefaultGridLayoutConfigurationProperty struct {
	// `CfnTemplate.DefaultGridLayoutConfigurationProperty.CanvasSizeOptions`.
	CanvasSizeOptions interface{} `field:"required" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

