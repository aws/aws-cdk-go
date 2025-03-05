package awsquicksight


// The label configuration of a reference line.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineLabelConfigurationProperty := &ReferenceLineLabelConfigurationProperty{
//   	CustomLabelConfiguration: &ReferenceLineCustomLabelConfigurationProperty{
//   		CustomLabel: jsii.String("customLabel"),
//   	},
//   	FontColor: jsii.String("fontColor"),
//   	FontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontSize: &FontSizeProperty{
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	HorizontalPosition: jsii.String("horizontalPosition"),
//   	ValueLabelConfiguration: &ReferenceLineValueLabelConfigurationProperty{
//   		FormatConfiguration: &NumericFormatConfigurationProperty{
//   			CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   				Symbol: jsii.String("symbol"),
//   			},
//   			NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				NumberScale: jsii.String("numberScale"),
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   			PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   				DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   					DecimalPlaces: jsii.Number(123),
//   				},
//   				NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   					DisplayMode: jsii.String("displayMode"),
//   				},
//   				NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   					NullString: jsii.String("nullString"),
//   				},
//   				Prefix: jsii.String("prefix"),
//   				SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   					DecimalSeparator: jsii.String("decimalSeparator"),
//   					ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   						Symbol: jsii.String("symbol"),
//   						Visibility: jsii.String("visibility"),
//   					},
//   				},
//   				Suffix: jsii.String("suffix"),
//   			},
//   		},
//   		RelativePosition: jsii.String("relativePosition"),
//   	},
//   	VerticalPosition: jsii.String("verticalPosition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html
//
type CfnTemplate_ReferenceLineLabelConfigurationProperty struct {
	// The custom label configuration of the label in a reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-customlabelconfiguration
	//
	CustomLabelConfiguration interface{} `field:"optional" json:"customLabelConfiguration" yaml:"customLabelConfiguration"`
	// The font color configuration of the label in a reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-fontcolor
	//
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// The font configuration of the label in a reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-fontconfiguration
	//
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
	// The horizontal position configuration of the label in a reference line. Choose one of the following options:.
	//
	// - `LEFT`
	// - `CENTER`
	// - `RIGHT`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-horizontalposition
	//
	HorizontalPosition *string `field:"optional" json:"horizontalPosition" yaml:"horizontalPosition"`
	// The value label configuration of the label in a reference line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-valuelabelconfiguration
	//
	ValueLabelConfiguration interface{} `field:"optional" json:"valueLabelConfiguration" yaml:"valueLabelConfiguration"`
	// The vertical position configuration of the label in a reference line. Choose one of the following options:.
	//
	// - `ABOVE`
	// - `BELOW`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinelabelconfiguration.html#cfn-quicksight-template-referencelinelabelconfiguration-verticalposition
	//
	VerticalPosition *string `field:"optional" json:"verticalPosition" yaml:"verticalPosition"`
}

