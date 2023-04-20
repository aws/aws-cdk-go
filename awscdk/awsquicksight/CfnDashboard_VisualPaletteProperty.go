package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualPaletteProperty := &VisualPaletteProperty{
//   	ChartColor: jsii.String("chartColor"),
//   	ColorMap: []interface{}{
//   		&DataPathColorProperty{
//   			Color: jsii.String("color"),
//   			Element: &DataPathValueProperty{
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   			},
//
//   			// the properties below are optional
//   			TimeGranularity: jsii.String("timeGranularity"),
//   		},
//   	},
//   }
//
type CfnDashboard_VisualPaletteProperty struct {
	// `CfnDashboard.VisualPaletteProperty.ChartColor`.
	ChartColor *string `field:"optional" json:"chartColor" yaml:"chartColor"`
	// `CfnDashboard.VisualPaletteProperty.ColorMap`.
	ColorMap interface{} `field:"optional" json:"colorMap" yaml:"colorMap"`
}

