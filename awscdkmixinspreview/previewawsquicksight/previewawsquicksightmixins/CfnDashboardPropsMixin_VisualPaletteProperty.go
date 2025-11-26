package previewawsquicksightmixins


// The visual display options for the visual palette.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   visualPaletteProperty := &VisualPaletteProperty{
//   	ChartColor: jsii.String("chartColor"),
//   	ColorMap: []interface{}{
//   		&DataPathColorProperty{
//   			Color: jsii.String("color"),
//   			Element: &DataPathValueProperty{
//   				DataPathType: &DataPathTypeProperty{
//   					PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   			},
//   			TimeGranularity: jsii.String("timeGranularity"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualpalette.html
//
type CfnDashboardPropsMixin_VisualPaletteProperty struct {
	// The chart color options for the visual palette.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualpalette.html#cfn-quicksight-dashboard-visualpalette-chartcolor
	//
	ChartColor *string `field:"optional" json:"chartColor" yaml:"chartColor"`
	// The color map options for the visual palette.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualpalette.html#cfn-quicksight-dashboard-visualpalette-colormap
	//
	ColorMap interface{} `field:"optional" json:"colorMap" yaml:"colorMap"`
}

