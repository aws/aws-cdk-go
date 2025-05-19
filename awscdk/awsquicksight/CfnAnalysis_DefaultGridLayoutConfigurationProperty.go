package awsquicksight


// The options that determine the default settings for a grid layout configuration.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultgridlayoutconfiguration.html
//
type CfnAnalysis_DefaultGridLayoutConfigurationProperty struct {
	// Determines the screen canvas size options for a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultgridlayoutconfiguration.html#cfn-quicksight-analysis-defaultgridlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"required" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

