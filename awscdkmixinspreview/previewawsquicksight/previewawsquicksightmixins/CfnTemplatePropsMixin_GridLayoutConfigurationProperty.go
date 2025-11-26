package previewawsquicksightmixins


// The configuration for a grid layout. Also called a tiled layout.
//
// Visuals snap to a grid with standard spacing and alignment. Dashboards are displayed as designed, with options to fit to screen or view at actual size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gridLayoutConfigurationProperty := &GridLayoutConfigurationProperty{
//   	CanvasSizeOptions: &GridLayoutCanvasSizeOptionsProperty{
//   		ScreenCanvasSizeOptions: &GridLayoutScreenCanvasSizeOptionsProperty{
//   			OptimizedViewPortWidth: jsii.String("optimizedViewPortWidth"),
//   			ResizeOption: jsii.String("resizeOption"),
//   		},
//   	},
//   	Elements: []interface{}{
//   		&GridLayoutElementProperty{
//   			ColumnIndex: jsii.Number(123),
//   			ColumnSpan: jsii.Number(123),
//   			ElementId: jsii.String("elementId"),
//   			ElementType: jsii.String("elementType"),
//   			RowIndex: jsii.Number(123),
//   			RowSpan: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutconfiguration.html
//
type CfnTemplatePropsMixin_GridLayoutConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutconfiguration.html#cfn-quicksight-template-gridlayoutconfiguration-canvassizeoptions
	//
	CanvasSizeOptions interface{} `field:"optional" json:"canvasSizeOptions" yaml:"canvasSizeOptions"`
	// The elements that are included in a grid layout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gridlayoutconfiguration.html#cfn-quicksight-template-gridlayoutconfiguration-elements
	//
	Elements interface{} `field:"optional" json:"elements" yaml:"elements"`
}

