package awsquicksight


// The options that determine the number display format configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberDisplayFormatConfigurationProperty := &NumberDisplayFormatConfigurationProperty{
//   	DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   		DecimalPlaces: jsii.Number(123),
//   	},
//   	NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   		DisplayMode: jsii.String("displayMode"),
//   	},
//   	NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   		NullString: jsii.String("nullString"),
//   	},
//   	NumberScale: jsii.String("numberScale"),
//   	Prefix: jsii.String("prefix"),
//   	SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   		DecimalSeparator: jsii.String("decimalSeparator"),
//   		ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   			GroupingStyle: jsii.String("groupingStyle"),
//   			Symbol: jsii.String("symbol"),
//   			Visibility: jsii.String("visibility"),
//   		},
//   	},
//   	Suffix: jsii.String("suffix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html
//
type CfnDashboard_NumberDisplayFormatConfigurationProperty struct {
	// The option that determines the decimal places configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-decimalplacesconfiguration
	//
	DecimalPlacesConfiguration interface{} `field:"optional" json:"decimalPlacesConfiguration" yaml:"decimalPlacesConfiguration"`
	// The options that determine the negative value configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-negativevalueconfiguration
	//
	NegativeValueConfiguration interface{} `field:"optional" json:"negativeValueConfiguration" yaml:"negativeValueConfiguration"`
	// The options that determine the null value format configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-nullvalueformatconfiguration
	//
	NullValueFormatConfiguration interface{} `field:"optional" json:"nullValueFormatConfiguration" yaml:"nullValueFormatConfiguration"`
	// Determines the number scale value of the number format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-numberscale
	//
	NumberScale *string `field:"optional" json:"numberScale" yaml:"numberScale"`
	// Determines the prefix value of the number format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The options that determine the numeric separator configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-separatorconfiguration
	//
	SeparatorConfiguration interface{} `field:"optional" json:"separatorConfiguration" yaml:"separatorConfiguration"`
	// Determines the suffix value of the number format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

