package interfacesawsconfig


// A reference to a OrganizationConfigRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationConfigRuleReference := &OrganizationConfigRuleReference{
//   	OrganizationConfigRuleArn: jsii.String("organizationConfigRuleArn"),
//   	OrganizationConfigRuleName: jsii.String("organizationConfigRuleName"),
//   }
//
type OrganizationConfigRuleReference struct {
	// The ARN of the OrganizationConfigRule resource.
	OrganizationConfigRuleArn *string `field:"required" json:"organizationConfigRuleArn" yaml:"organizationConfigRuleArn"`
	// The OrganizationConfigRuleName of the OrganizationConfigRule resource.
	OrganizationConfigRuleName *string `field:"required" json:"organizationConfigRuleName" yaml:"organizationConfigRuleName"`
}

