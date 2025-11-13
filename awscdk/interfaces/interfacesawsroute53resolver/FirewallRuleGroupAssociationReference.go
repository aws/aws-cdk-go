package interfacesawsroute53resolver


// A reference to a FirewallRuleGroupAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallRuleGroupAssociationReference := &FirewallRuleGroupAssociationReference{
//   	FirewallRuleGroupAssociationArn: jsii.String("firewallRuleGroupAssociationArn"),
//   	FirewallRuleGroupAssociationId: jsii.String("firewallRuleGroupAssociationId"),
//   }
//
type FirewallRuleGroupAssociationReference struct {
	// The ARN of the FirewallRuleGroupAssociation resource.
	FirewallRuleGroupAssociationArn *string `field:"required" json:"firewallRuleGroupAssociationArn" yaml:"firewallRuleGroupAssociationArn"`
	// The Id of the FirewallRuleGroupAssociation resource.
	FirewallRuleGroupAssociationId *string `field:"required" json:"firewallRuleGroupAssociationId" yaml:"firewallRuleGroupAssociationId"`
}

