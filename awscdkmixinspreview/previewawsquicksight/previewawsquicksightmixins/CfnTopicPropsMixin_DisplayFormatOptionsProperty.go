package previewawsquicksightmixins


// A structure that represents additional options for display formatting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html
//
type CfnTopicPropsMixin_DisplayFormatOptionsProperty struct {
	// Determines the blank cell format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-blankcellformat
	//
	BlankCellFormat *string `field:"optional" json:"blankCellFormat" yaml:"blankCellFormat"`
	// The currency symbol, such as `USD` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-currencysymbol
	//
	CurrencySymbol *string `field:"optional" json:"currencySymbol" yaml:"currencySymbol"`
	// Determines the `DateTime` format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-dateformat
	//
	DateFormat *string `field:"optional" json:"dateFormat" yaml:"dateFormat"`
	// Determines the decimal separator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-decimalseparator
	//
	DecimalSeparator *string `field:"optional" json:"decimalSeparator" yaml:"decimalSeparator"`
	// Determines the number of fraction digits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-fractiondigits
	//
	// Default: - 0.
	//
	FractionDigits *float64 `field:"optional" json:"fractionDigits" yaml:"fractionDigits"`
	// Determines the grouping separator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-groupingseparator
	//
	GroupingSeparator *string `field:"optional" json:"groupingSeparator" yaml:"groupingSeparator"`
	// The negative format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-negativeformat
	//
	NegativeFormat interface{} `field:"optional" json:"negativeFormat" yaml:"negativeFormat"`
	// The prefix value for a display format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The suffix value for a display format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-suffix
	//
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
	// The unit scaler.
	//
	// Valid values for this structure are: `NONE` , `AUTO` , `THOUSANDS` , `MILLIONS` , `BILLIONS` , and `TRILLIONS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-unitscaler
	//
	UnitScaler *string `field:"optional" json:"unitScaler" yaml:"unitScaler"`
	// A Boolean value that indicates whether to use blank cell format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-useblankcellformat
	//
	// Default: - false.
	//
	UseBlankCellFormat interface{} `field:"optional" json:"useBlankCellFormat" yaml:"useBlankCellFormat"`
	// A Boolean value that indicates whether to use grouping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-displayformatoptions.html#cfn-quicksight-topic-displayformatoptions-usegrouping
	//
	// Default: - false.
	//
	UseGrouping interface{} `field:"optional" json:"useGrouping" yaml:"useGrouping"`
}

