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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html#cfn-securityhub-organizationconfiguration-autoenable
	//
	AutoEnable interface{} `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Whether to automatically enable Security Hub default standards in new member accounts when they join the organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html#cfn-securityhub-organizationconfiguration-autoenablestandards
	//
	AutoEnableStandards *string `field:"optional" json:"autoEnableStandards" yaml:"autoEnableStandards"`
	// Indicates whether the organization uses local or central configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-organizationconfiguration.html#cfn-securityhub-organizationconfiguration-configurationtype
	//
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
}

