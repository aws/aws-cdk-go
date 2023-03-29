package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericSeparatorConfigurationProperty := &NumericSeparatorConfigurationProperty{
//   	DecimalSeparator: jsii.String("decimalSeparator"),
//   	ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   		Symbol: jsii.String("symbol"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnDashboard_NumericSeparatorConfigurationProperty struct {
	// `CfnDashboard.NumericSeparatorConfigurationProperty.DecimalSeparator`.
	DecimalSeparator *string `field:"optional" json:"decimalSeparator" yaml:"decimalSeparator"`
	// `CfnDashboard.NumericSeparatorConfigurationProperty.ThousandsSeparator`.
	ThousandsSeparator interface{} `field:"optional" json:"thousandsSeparator" yaml:"thousandsSeparator"`
}

