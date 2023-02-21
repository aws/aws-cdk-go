package awsrds


// The `OptionSetting` property type specifies the value for an option within an `OptionSetting` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionSettingProperty := &optionSettingProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnOptionGroup_OptionSettingProperty struct {
	// The name of the option that has settings that you can set.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The current value of the option setting.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

