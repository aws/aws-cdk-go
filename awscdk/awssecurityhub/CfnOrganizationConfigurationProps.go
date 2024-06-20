package awssecurityhub


// Properties for defining a `CfnOrganizationConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationConfigurationProps := &CfnOrganizationConfigurationProps{
//   	AutoEnable: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AutoEnableStandards: jsii.String("autoEnableStandards"),
//   	ConfigurationType: jsii.String("configurationType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html
//
type CfnOrganizationConfigurationProps struct {
	// Whether to automatically enable Security Hub in new member accounts when they join the organization.
	//
	// If set to `true` , then Security Hub is automatically enabled in new accounts. If set to `false` , then Security Hub isn't enabled in new accounts automatically. The default value is `false` .
	//
	// If the `ConfigurationType` of your organization is set to `CENTRAL` , then this field is set to `false` and can't be changed in the home Region and linked Regions. However, in that case, the delegated administrator can create a configuration policy in which Security Hub is enabled and associate the policy with new organization accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html#cfn-securityhub-organizationconfiguration-autoenable
	//
	AutoEnable interface{} `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Whether to automatically enable Security Hub [default standards](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-enable-disable.html) in new member accounts when they join the organization.
	//
	// The default value of this parameter is equal to `DEFAULT` .
	//
	// If equal to `DEFAULT` , then Security Hub default standards are automatically enabled for new member accounts. If equal to `NONE` , then default standards are not automatically enabled for new member accounts.
	//
	// If the `ConfigurationType` of your organization is set to `CENTRAL` , then this field is set to `NONE` and can't be changed in the home Region and linked Regions. However, in that case, the delegated administrator can create a configuration policy in which specific security standards are enabled and associate the policy with new organization accounts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html#cfn-securityhub-organizationconfiguration-autoenablestandards
	//
	AutoEnableStandards *string `field:"optional" json:"autoEnableStandards" yaml:"autoEnableStandards"`
	// Indicates whether the organization uses local or central configuration.
	//
	// If you use local configuration, the Security Hub delegated administrator can set `AutoEnable` to `true` and `AutoEnableStandards` to `DEFAULT` . This automatically enables Security Hub and default security standards in new organization accounts. These new account settings must be set separately in each AWS Region , and settings may be different in each Region.
	//
	// If you use central configuration, the delegated administrator can create configuration policies. Configuration policies can be used to configure Security Hub, security standards, and security controls in multiple accounts and Regions. If you want new organization accounts to use a specific configuration, you can create a configuration policy and associate it with the root or specific organizational units (OUs). New accounts will inherit the policy from the root or their assigned OU.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html#cfn-securityhub-organizationconfiguration-configurationtype
	//
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
}

