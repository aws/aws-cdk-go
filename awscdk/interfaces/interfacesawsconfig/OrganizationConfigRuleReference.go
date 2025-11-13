package interfacesawsconfig


// A reference to a OrganizationConfigRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationConfigRuleReference := &OrganizationConfigRuleReference{
//   	OrganizationConfigRuleId: jsii.String("organizationConfigRuleId"),
//   }
//
type OrganizationConfigRuleReference struct {
	// The Id of the OrganizationConfigRule resource.
	OrganizationConfigRuleId *string `field:"required" json:"organizationConfigRuleId" yaml:"organizationConfigRuleId"`
}

