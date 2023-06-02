package awsquicksight


// A structure that represents additional options for display formatting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   displayFormatOptionsProperty := &DisplayFormatOptionsProperty{
//   	BlankCellFormat: jsii.String("blankCellFormat"),
//   	CurrencySymbol: jsii.String("currencySymbol"),
//   	DateFormat: jsii.String("dateFormat"),
//   	DecimalSeparator: jsii.String("decimalSeparator"),
//   	FractionDigits: jsii.Number(123),
//   	GroupingSeparator: jsii.String("groupingSeparator"),
//   	NegativeFormat: &NegativeFormatProperty{
//   		Prefix: jsii.String("prefix"),
//   		Suffix: jsii.String("suffix"),
//   	},
//   	Prefix: jsii.String("prefix"),
//   	Suffix: jsii.String("suffix"),
//   	UnitScaler: jsii.String("unitScaler"),
//   	UseBlankCellFormat: jsii.Boolean(false),
//   	UseGrouping: jsii.Boolean(false),
//   }
//
type CfnTopic_DisplayFormatOptionsProperty struct {
	// Determines the blank cell format.
	BlankCellFormat *string `field:"optional" json:"blankCellFormat" yaml:"blankCellFormat"`
	// The currency symbol, such as `USD` .
	CurrencySymbol *string `field:"optional" json:"currencySymbol" yaml:"currencySymbol"`
	// Determines the `DateTime` format.
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Determines the decimal separator.
	DecimalSeparator *string `field:"optional" json:"decimalSeparator" yaml:"decimalSeparator"`
	// Determines the number of fraction digits.
	FractionDigits *float64 `field:"optional" json:"fractionDigits" yaml:"fractionDigits"`
	// Determines the grouping separator.
	GroupingSeparator *string `field:"optional" json:"groupingSeparator" yaml:"groupingSeparator"`
	// The negative format.
	NegativeFormat interface{} `field:"optional" json:"negativeFormat" yaml:"negativeFormat"`
	// The prefix value for a display format.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The suffix value for a display format.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// The unit scaler.
	//
	// Valid values for this structure are: `NONE` , `AUTO` , `THOUSANDS` , `MILLIONS` , `BILLIONS` , and `TRILLIONS` .
	UnitScaler *string `field:"optional" json:"unitScaler" yaml:"unitScaler"`
	// A Boolean value that indicates whether to use blank cell format.
	UseBlankCellFormat interface{} `field:"optional" json:"useBlankCellFormat" yaml:"useBlankCellFormat"`
	// A Boolean value that indicates whether to use grouping.
	UseGrouping interface{} `field:"optional" json:"useGrouping" yaml:"useGrouping"`
}

