package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integerValueWhenUnsetConfigurationProperty := &IntegerValueWhenUnsetConfigurationProperty{
//   	CustomValue: jsii.Number(123),
//   	ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   }
//
type CfnDashboard_IntegerValueWhenUnsetConfigurationProperty struct {
	// `CfnDashboard.IntegerValueWhenUnsetConfigurationProperty.CustomValue`.
	CustomValue *float64 `field:"optional" json:"customValue" yaml:"customValue"`
	// `CfnDashboard.IntegerValueWhenUnsetConfigurationProperty.ValueWhenUnsetOption`.
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

