package awsquicksight


// The theme colors that apply to UI and to charts, excluding data colors.
//
// The colors description is a hexadecimal color code that consists of six alphanumerical characters, prefixed with `#` , for example #37BFF5. For more information, see [Using Themes in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html) in the *Amazon QuickSight User Guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uIColorPaletteProperty := &uIColorPaletteProperty{
//   	accent: jsii.String("accent"),
//   	accentForeground: jsii.String("accentForeground"),
//   	danger: jsii.String("danger"),
//   	dangerForeground: jsii.String("dangerForeground"),
//   	dimension: jsii.String("dimension"),
//   	dimensionForeground: jsii.String("dimensionForeground"),
//   	measure: jsii.String("measure"),
//   	measureForeground: jsii.String("measureForeground"),
//   	primaryBackground: jsii.String("primaryBackground"),
//   	primaryForeground: jsii.String("primaryForeground"),
//   	secondaryBackground: jsii.String("secondaryBackground"),
//   	secondaryForeground: jsii.String("secondaryForeground"),
//   	success: jsii.String("success"),
//   	successForeground: jsii.String("successForeground"),
//   	warning: jsii.String("warning"),
//   	warningForeground: jsii.String("warningForeground"),
//   }
//
type CfnTheme_UIColorPaletteProperty struct {
	// This color is that applies to selected states and buttons.
	Accent *string `field:"optional" json:"accent" yaml:"accent"`
	// The foreground color that applies to any text or other elements that appear over the accent color.
	AccentForeground *string `field:"optional" json:"accentForeground" yaml:"accentForeground"`
	// The color that applies to error messages.
	Danger *string `field:"optional" json:"danger" yaml:"danger"`
	// The foreground color that applies to any text or other elements that appear over the error color.
	DangerForeground *string `field:"optional" json:"dangerForeground" yaml:"dangerForeground"`
	// The color that applies to the names of fields that are identified as dimensions.
	Dimension *string `field:"optional" json:"dimension" yaml:"dimension"`
	// The foreground color that applies to any text or other elements that appear over the dimension color.
	DimensionForeground *string `field:"optional" json:"dimensionForeground" yaml:"dimensionForeground"`
	// The color that applies to the names of fields that are identified as measures.
	Measure *string `field:"optional" json:"measure" yaml:"measure"`
	// The foreground color that applies to any text or other elements that appear over the measure color.
	MeasureForeground *string `field:"optional" json:"measureForeground" yaml:"measureForeground"`
	// The background color that applies to visuals and other high emphasis UI.
	PrimaryBackground *string `field:"optional" json:"primaryBackground" yaml:"primaryBackground"`
	// The color of text and other foreground elements that appear over the primary background regions, such as grid lines, borders, table banding, icons, and so on.
	PrimaryForeground *string `field:"optional" json:"primaryForeground" yaml:"primaryForeground"`
	// The background color that applies to the sheet background and sheet controls.
	SecondaryBackground *string `field:"optional" json:"secondaryBackground" yaml:"secondaryBackground"`
	// The foreground color that applies to any sheet title, sheet control text, or UI that appears over the secondary background.
	SecondaryForeground *string `field:"optional" json:"secondaryForeground" yaml:"secondaryForeground"`
	// The color that applies to success messages, for example the check mark for a successful download.
	Success *string `field:"optional" json:"success" yaml:"success"`
	// The foreground color that applies to any text or other elements that appear over the success color.
	SuccessForeground *string `field:"optional" json:"successForeground" yaml:"successForeground"`
	// This color that applies to warning and informational messages.
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
	// The foreground color that applies to any text or other elements that appear over the warning color.
	WarningForeground *string `field:"optional" json:"warningForeground" yaml:"warningForeground"`
}

