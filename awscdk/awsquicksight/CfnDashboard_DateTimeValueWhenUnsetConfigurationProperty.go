package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeValueWhenUnsetConfigurationProperty := &DateTimeValueWhenUnsetConfigurationProperty{
//   	CustomValue: jsii.String("customValue"),
//   	ValueWhenUnsetOption: jsii.String("valueWhenUnsetOption"),
//   }
//
type CfnDashboard_DateTimeValueWhenUnsetConfigurationProperty struct {
	// `CfnDashboard.DateTimeValueWhenUnsetConfigurationProperty.CustomValue`.
	CustomValue *string `field:"optional" json:"customValue" yaml:"customValue"`
	// `CfnDashboard.DateTimeValueWhenUnsetConfigurationProperty.ValueWhenUnsetOption`.
	ValueWhenUnsetOption *string `field:"optional" json:"valueWhenUnsetOption" yaml:"valueWhenUnsetOption"`
}

