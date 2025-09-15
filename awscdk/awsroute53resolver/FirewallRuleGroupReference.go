package awsroute53resolver


// A reference to a FirewallRuleGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallRuleGroupReference := &FirewallRuleGroupReference{
//   	FirewallRuleGroupArn: jsii.String("firewallRuleGroupArn"),
//   	FirewallRuleGroupId: jsii.String("firewallRuleGroupId"),
//   }
//
type FirewallRuleGroupReference struct {
	// The ARN of the FirewallRuleGroup resource.
	FirewallRuleGroupArn *string `field:"required" json:"firewallRuleGroupArn" yaml:"firewallRuleGroupArn"`
	// The Id of the FirewallRuleGroup resource.
	FirewallRuleGroupId *string `field:"required" json:"firewallRuleGroupId" yaml:"firewallRuleGroupId"`
}

