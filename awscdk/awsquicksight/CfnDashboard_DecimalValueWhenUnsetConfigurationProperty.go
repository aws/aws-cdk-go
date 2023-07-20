package awsquicksight


// The configuration that defines the default value of a `Decimal` parameter when a value has not been set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalValueWhenUnsetConfigurationProperty := &DecimalValueWhenUnsetConfigurationProperty{
//   	CustomValue: jsii.Number(123),
//   	ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalvaluewhenunsetconfiguration.html
//
type CfnDashboard_DecimalValueWhenUnsetConfigurationProperty struct {
	// A custom value that's used when the value of a parameter isn't set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalvaluewhenunsetconfiguration.html#cfn-quicksight-dashboard-decimalvaluewhenunsetconfiguration-customvalue
	//
	CustomValue *float64 `field:"optional" json:"customValue" yaml:"customValue"`
	// The built-in options for default values. The value can be one of the following:.
	//
	// - `RECOMMENDED` : The recommended value.
	// - `NULL` : The `NULL` value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalvaluewhenunsetconfiguration.html#cfn-quicksight-dashboard-decimalvaluewhenunsetconfiguration-valuewhenunsetoption
	//
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

