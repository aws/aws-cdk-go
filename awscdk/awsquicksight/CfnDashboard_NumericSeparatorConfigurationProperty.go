package awsquicksight


// The options that determine the numeric separator configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericSeparatorConfigurationProperty := &NumericSeparatorConfigurationProperty{
//   	DecimalSeparator: jsii.String("decimalSeparator"),
//   	ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   		GroupingStyle: jsii.String("groupingStyle"),
//   		Symbol: jsii.String("symbol"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericseparatorconfiguration.html
//
type CfnDashboard_NumericSeparatorConfigurationProperty struct {
	// Determines the decimal separator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericseparatorconfiguration.html#cfn-quicksight-dashboard-numericseparatorconfiguration-decimalseparator
	//
	DecimalSeparator *string `field:"optional" json:"decimalSeparator" yaml:"decimalSeparator"`
	// The options that determine the thousands separator configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numericseparatorconfiguration.html#cfn-quicksight-dashboard-numericseparatorconfiguration-thousandsseparator
	//
	ThousandsSeparator interface{} `field:"optional" json:"thousandsSeparator" yaml:"thousandsSeparator"`
}

