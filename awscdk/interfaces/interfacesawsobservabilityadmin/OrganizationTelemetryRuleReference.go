package interfacesawsobservabilityadmin


// A reference to a OrganizationTelemetryRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationTelemetryRuleReference := &OrganizationTelemetryRuleReference{
//   	RuleArn: jsii.String("ruleArn"),
//   }
//
type OrganizationTelemetryRuleReference struct {
	// The RuleArn of the OrganizationTelemetryRule resource.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
}

