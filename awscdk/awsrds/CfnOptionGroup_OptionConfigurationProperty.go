package awsrds


// The `OptionConfiguration` property type specifies an individual option, and its settings, within an `AWS::RDS::OptionGroup` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionConfigurationProperty := &OptionConfigurationProperty{
//   	OptionName: jsii.String("optionName"),
//
//   	// the properties below are optional
//   	DbSecurityGroupMemberships: []*string{
//   		jsii.String("dbSecurityGroupMemberships"),
//   	},
//   	OptionSettings: []interface{}{
//   		&OptionSettingProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	OptionVersion: jsii.String("optionVersion"),
//   	Port: jsii.Number(123),
//   	VpcSecurityGroupMemberships: []*string{
//   		jsii.String("vpcSecurityGroupMemberships"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html
//
type CfnOptionGroup_OptionConfigurationProperty struct {
	// The configuration of options to include in a group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html#cfn-rds-optiongroup-optionconfiguration-optionname
	//
	OptionName *string `field:"required" json:"optionName" yaml:"optionName"`
	// A list of DBSecurityGroupMembership name strings used for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html#cfn-rds-optiongroup-optionconfiguration-dbsecuritygroupmemberships
	//
	DbSecurityGroupMemberships *[]*string `field:"optional" json:"dbSecurityGroupMemberships" yaml:"dbSecurityGroupMemberships"`
	// The option settings to include in an option group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html#cfn-rds-optiongroup-optionconfiguration-optionsettings
	//
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The version for the option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html#cfn-rds-optiongroup-optionconfiguration-optionversion
	//
	OptionVersion *string `field:"optional" json:"optionVersion" yaml:"optionVersion"`
	// The optional port for the option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html#cfn-rds-optiongroup-optionconfiguration-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of VpcSecurityGroupMembership name strings used for this option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-optiongroup-optionconfiguration.html#cfn-rds-optiongroup-optionconfiguration-vpcsecuritygroupmemberships
	//
	VpcSecurityGroupMemberships *[]*string `field:"optional" json:"vpcSecurityGroupMemberships" yaml:"vpcSecurityGroupMemberships"`
}

