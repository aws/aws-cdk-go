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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultinteractivelayoutconfiguration.html
//
type CfnAnalysis_DefaultInteractiveLayoutConfigurationProperty struct {
	// The options that determine the default settings of a free-form layout configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultinteractivelayoutconfiguration.html#cfn-quicksight-analysis-defaultinteractivelayoutconfiguration-freeform
	//
	FreeForm interface{} `field:"optional" json:"freeForm" yaml:"freeForm"`
	// The options that determine the default settings for a grid layout configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultinteractivelayoutconfiguration.html#cfn-quicksight-analysis-defaultinteractivelayoutconfiguration-grid
	//
	Grid interface{} `field:"optional" json:"grid" yaml:"grid"`
}

