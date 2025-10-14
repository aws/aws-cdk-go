package awsobservabilityadmin


// A reference to a OrganizationCentralizationRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationCentralizationRuleReference := &OrganizationCentralizationRuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type OrganizationCentralizationRuleReference struct {
	// The RuleArn of the OrganizationCentralizationRule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

