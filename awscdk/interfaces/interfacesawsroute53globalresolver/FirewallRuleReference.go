package interfacesawsroute53globalresolver


// A reference to a FirewallRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallRuleReference := &FirewallRuleReference{
//   	FirewallRuleId: jsii.String("firewallRuleId"),
//   }
//
type FirewallRuleReference struct {
	// The FirewallRuleId of the FirewallRule resource.
	FirewallRuleId *string `field:"required" json:"firewallRuleId" yaml:"firewallRuleId"`
}

