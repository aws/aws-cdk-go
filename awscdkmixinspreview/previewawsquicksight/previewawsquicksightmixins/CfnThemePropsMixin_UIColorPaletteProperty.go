package previewawsquicksightmixins


// The theme colors that apply to UI and to charts, excluding data colors.
//
// The colors description is a hexadecimal color code that consists of six alphanumerical characters, prefixed with `#` , for example #37BFF5. For more information, see [Using Themes in Amazon Quick Suite](https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html) in the *Amazon Quick Suite User Guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uIColorPaletteProperty := &UIColorPaletteProperty{
//   	Accent: jsii.String("accent"),
//   	AccentForeground: jsii.String("accentForeground"),
//   	Danger: jsii.String("danger"),
//   	DangerForeground: jsii.String("dangerForeground"),
//   	Dimension: jsii.String("dimension"),
//   	DimensionForeground: jsii.String("dimensionForeground"),
//   	Measure: jsii.String("measure"),
//   	MeasureForeground: jsii.String("measureForeground"),
//   	PrimaryBackground: jsii.String("primaryBackground"),
//   	PrimaryForeground: jsii.String("primaryForeground"),
//   	SecondaryBackground: jsii.String("secondaryBackground"),
//   	SecondaryForeground: jsii.String("secondaryForeground"),
//   	Success: jsii.String("success"),
//   	SuccessForeground: jsii.String("successForeground"),
//   	Warning: jsii.String("warning"),
//   	WarningForeground: jsii.String("warningForeground"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html
//
type CfnThemePropsMixin_UIColorPaletteProperty struct {
	// This color is that applies to selected states and buttons.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-accent
	//
	Accent *string `field:"optional" json:"accent" yaml:"accent"`
	// The foreground color that applies to any text or other elements that appear over the accent color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-accentforeground
	//
	AccentForeground *string `field:"optional" json:"accentForeground" yaml:"accentForeground"`
	// The color that applies to error messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-danger
	//
	Danger *string `field:"optional" json:"danger" yaml:"danger"`
	// The foreground color that applies to any text or other elements that appear over the error color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-dangerforeground
	//
	DangerForeground *string `field:"optional" json:"dangerForeground" yaml:"dangerForeground"`
	// The color that applies to the names of fields that are identified as dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-dimension
	//
	Dimension *string `field:"optional" json:"dimension" yaml:"dimension"`
	// The foreground color that applies to any text or other elements that appear over the dimension color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-dimensionforeground
	//
	DimensionForeground *string `field:"optional" json:"dimensionForeground" yaml:"dimensionForeground"`
	// The color that applies to the names of fields that are identified as measures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-measure
	//
	Measure *string `field:"optional" json:"measure" yaml:"measure"`
	// The foreground color that applies to any text or other elements that appear over the measure color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-measureforeground
	//
	MeasureForeground *string `field:"optional" json:"measureForeground" yaml:"measureForeground"`
	// The background color that applies to visuals and other high emphasis UI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-primarybackground
	//
	PrimaryBackground *string `field:"optional" json:"primaryBackground" yaml:"primaryBackground"`
	// The color of text and other foreground elements that appear over the primary background regions, such as grid lines, borders, table banding, icons, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-primaryforeground
	//
	PrimaryForeground *string `field:"optional" json:"primaryForeground" yaml:"primaryForeground"`
	// The background color that applies to the sheet background and sheet controls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-secondarybackground
	//
	SecondaryBackground *string `field:"optional" json:"secondaryBackground" yaml:"secondaryBackground"`
	// The foreground color that applies to any sheet title, sheet control text, or UI that appears over the secondary background.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-secondaryforeground
	//
	SecondaryForeground *string `field:"optional" json:"secondaryForeground" yaml:"secondaryForeground"`
	// The color that applies to success messages, for example the check mark for a successful download.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-success
	//
	Success *string `field:"optional" json:"success" yaml:"success"`
	// The foreground color that applies to any text or other elements that appear over the success color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-successforeground
	//
	SuccessForeground *string `field:"optional" json:"successForeground" yaml:"successForeground"`
	// This color that applies to warning and informational messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-warning
	//
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
	// The foreground color that applies to any text or other elements that appear over the warning color.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-uicolorpalette.html#cfn-quicksight-theme-uicolorpalette-warningforeground
	//
	WarningForeground *string `field:"optional" json:"warningForeground" yaml:"warningForeground"`
}

