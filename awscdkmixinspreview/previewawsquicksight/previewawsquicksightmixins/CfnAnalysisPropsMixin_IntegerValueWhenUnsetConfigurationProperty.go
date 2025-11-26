package previewawsquicksightmixins


// A parameter declaration for the `Integer` data type.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   integerValueWhenUnsetConfigurationProperty := &IntegerValueWhenUnsetConfigurationProperty{
//   	CustomValue: jsii.Number(123),
//   	ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-integervaluewhenunsetconfiguration.html
//
type CfnAnalysisPropsMixin_IntegerValueWhenUnsetConfigurationProperty struct {
	// A custom value that's used when the value of a parameter isn't set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-integervaluewhenunsetconfiguration.html#cfn-quicksight-analysis-integervaluewhenunsetconfiguration-customvalue
	//
	CustomValue *float64 `field:"optional" json:"customValue" yaml:"customValue"`
	// The built-in options for default values. The value can be one of the following:.
	//
	// - `RECOMMENDED` : The recommended value.
	// - `NULL` : The `NULL` value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-integervaluewhenunsetconfiguration.html#cfn-quicksight-analysis-integervaluewhenunsetconfiguration-valuewhenunsetoption
	//
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

