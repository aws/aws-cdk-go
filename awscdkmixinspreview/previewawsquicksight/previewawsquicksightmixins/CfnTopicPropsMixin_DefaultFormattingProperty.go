package previewawsquicksightmixins


// A structure that represents a default formatting definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultFormattingProperty := &DefaultFormattingProperty{
//   	DisplayFormat: jsii.String("displayFormat"),
//   	DisplayFormatOptions: &DisplayFormatOptionsProperty{
//   		BlankCellFormat: jsii.String("blankCellFormat"),
//   		CurrencySymbol: jsii.String("currencySymbol"),
//   		DateFormat: jsii.String("dateFormat"),
//   		DecimalSeparator: jsii.String("decimalSeparator"),
//   		FractionDigits: jsii.Number(123),
//   		GroupingSeparator: jsii.String("groupingSeparator"),
//   		NegativeFormat: &NegativeFormatProperty{
//   			Prefix: jsii.String("prefix"),
//   			Suffix: jsii.String("suffix"),
//   		},
//   		Prefix: jsii.String("prefix"),
//   		Suffix: jsii.String("suffix"),
//   		UnitScaler: jsii.String("unitScaler"),
//   		UseBlankCellFormat: jsii.Boolean(false),
//   		UseGrouping: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-defaultformatting.html
//
type CfnTopicPropsMixin_DefaultFormattingProperty struct {
	// The display format.
	//
	// Valid values for this structure are `AUTO` , `PERCENT` , `CURRENCY` , `NUMBER` , `DATE` , and `STRING` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-defaultformatting.html#cfn-quicksight-topic-defaultformatting-displayformat
	//
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// The additional options for display formatting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-defaultformatting.html#cfn-quicksight-topic-defaultformatting-displayformatoptions
	//
	DisplayFormatOptions interface{} `field:"optional" json:"displayFormatOptions" yaml:"displayFormatOptions"`
}

