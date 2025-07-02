package awsrds


// The `OptionSetting` property type specifies the value for an option within an `OptionSetting` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionSettingProperty := &OptionSettingProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionsetting.html
//
type CfnOptionGroup_OptionSettingProperty struct {
	// The name of the option that has settings that you can set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionsetting.html#cfn-rds-optiongroup-optionsetting-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The current value of the option setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionsetting.html#cfn-rds-optiongroup-optionsetting-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

