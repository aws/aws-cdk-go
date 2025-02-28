package awsquicksight


// The dimension type field with date type columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateDimensionFieldProperty := &DateDimensionFieldProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FieldId: jsii.String("fieldId"),
//
//   	// the properties below are optional
//   	DateGranularity: jsii.String("dateGranularity"),
//   	FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   		DateTimeFormat: jsii.String("dateTimeFormat"),
//   		NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   			NullString: jsii.String("nullString"),
//   		},
//   		NumericFormatConfiguration: &NumericFormatConfigurationProperty{
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
//   	},
//   	HierarchyId: jsii.String("hierarchyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html
//
type CfnDashboard_DateDimensionFieldProperty struct {
	// The column that is used in the `DateDimensionField` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The custom field ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The date granularity of the `DateDimensionField` . Choose one of the following options:.
	//
	// - `YEAR`
	// - `QUARTER`
	// - `MONTH`
	// - `WEEK`
	// - `DAY`
	// - `HOUR`
	// - `MINUTE`
	// - `SECOND`
	// - `MILLISECOND`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-dategranularity
	//
	DateGranularity *string `field:"optional" json:"dateGranularity" yaml:"dateGranularity"`
	// The format configuration of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-formatconfiguration
	//
	FormatConfiguration interface{} `field:"optional" json:"formatConfiguration" yaml:"formatConfiguration"`
	// The custom hierarchy ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datedimensionfield.html#cfn-quicksight-dashboard-datedimensionfield-hierarchyid
	//
	HierarchyId *string `field:"optional" json:"hierarchyId" yaml:"hierarchyId"`
}

