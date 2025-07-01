package awsquicksight


// The options that determine the default settings of a free-form layout configuration.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfreeformlayoutconfiguration.html
//
type CfnAnalysis_DefaultFreeFormLayoutConfigurationProperty struct {
	// Determines the screen canvas size options for a free-form layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultfreeformlayoutconfiguration.html#cfn-quicksight-analysis-defaultfreeformlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"required" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

