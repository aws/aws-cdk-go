package awsquicksight


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
	// `CfnAnalysis.DefaultInteractiveLayoutConfigurationProperty.FreeForm`.
	FreeForm interface{} `field:"optional" json:"freeForm" yaml:"freeForm"`
	// `CfnAnalysis.DefaultInteractiveLayoutConfigurationProperty.Grid`.
	Grid interface{} `field:"optional" json:"grid" yaml:"grid"`
}

