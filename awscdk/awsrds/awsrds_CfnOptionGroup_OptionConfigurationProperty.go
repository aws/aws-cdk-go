package awsrds


// The `OptionConfiguration` property type specifies an individual option, and its settings, within an `AWS::RDS::OptionGroup` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionConfigurationProperty := &optionConfigurationProperty{
//   	optionName: jsii.String("optionName"),
//
//   	// the properties below are optional
//   	dbSecurityGroupMemberships: []*string{
//   		jsii.String("dbSecurityGroupMemberships"),
//   	},
//   	optionSettings: []interface{}{
//   		&optionSettingProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	optionVersion: jsii.String("optionVersion"),
//   	port: jsii.Number(123),
//   	vpcSecurityGroupMemberships: []*string{
//   		jsii.String("vpcSecurityGroupMemberships"),
//   	},
//   }
//
type CfnOptionGroup_OptionConfigurationProperty struct {
	// The configuration of options to include in a group.
	OptionName *string `field:"required" json:"optionName" yaml:"optionName"`
	// A list of DBSecurityGroupMembership name strings used for this option.
	DbSecurityGroupMemberships *[]*string `field:"optional" json:"dbSecurityGroupMemberships" yaml:"dbSecurityGroupMemberships"`
	// The option settings to include in an option group.
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The version for the option.
	OptionVersion *string `field:"optional" json:"optionVersion" yaml:"optionVersion"`
	// The optional port for the option.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of VpcSecurityGroupMembership name strings used for this option.
	VpcSecurityGroupMemberships *[]*string `field:"optional" json:"vpcSecurityGroupMemberships" yaml:"vpcSecurityGroupMemberships"`
}

