package awsquicksight


// The options that determine the default settings for interactive layout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultInteractiveLayoutConfigurationProperty := &DefaultInteractiveLayoutConfigurationProperty{
//   	FreeForm: &DefaultFreeFormLayoutConfigurationProperty{
//   		CanvasSizeOptions: &FreeFormLayoutCanvasSizeOptionsProperty{
//   			ScreenCanvasSizeOptions: &FreeFormLayoutScreenCanvasSizeOptionsProperty{
//   				OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   			},
//   		},
//   	},
//   	Grid: &DefaultGridLayoutConfigurationProperty{
//   		CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   			ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   				ResizeOption: jsii.String("resizeOption"),
//
//   				// the properties below are optional
//   				OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_DefaultInteractiveLayoutConfigurationProperty struct {
	// The options that determine the default settings of a free-form layout configuration.
	FreeForm interface{} `field:"optional" json:"freeForm" yaml:"freeForm"`
	// The options that determine the default settings for a grid layout configuration.
	Grid interface{} `field:"optional" json:"grid" yaml:"grid"`
}

