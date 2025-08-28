package awsquicksight


// The configuration for a grid layout. Also called a tiled layout.
//
// Visuals snap to a grid with standard spacing and alignment. Dashboards are displayed as designed, with options to fit to screen or view at actual size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gridLayoutConfigurationProperty := &GridLayoutConfigurationProperty{
//   	Elements: []interface{}{
//   		&GridLayoutElementProperty{
//   			ColumnSpan: jsii.Number(123),
//   			ElementId: jsii.String("elementId"),
//   			ElementType: jsii.String("elementType"),
//   			RowSpan: jsii.Number(123),
//
//   			// the properties below are optional
//   			ColumnIndex: jsii.Number(123),
//   			RowIndex: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutconfiguration.html
//
type CfnAnalysis_GridLayoutConfigurationProperty struct {
	// The elements that are included in a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutconfiguration.html#cfn-quicksight-analysis-gridlayoutconfiguration-elements
	//
	Elements interface{} `field:"required" json:"elements" yaml:"elements"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-gridlayoutconfiguration.html#cfn-quicksight-analysis-gridlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"optional" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
}

