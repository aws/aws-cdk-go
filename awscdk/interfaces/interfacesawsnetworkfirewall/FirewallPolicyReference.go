package interfacesawsnetworkfirewall


// A reference to a FirewallPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallPolicyReference := &FirewallPolicyReference{
//   	FirewallPolicyArn: jsii.String("firewallPolicyArn"),
//   }
//
type FirewallPolicyReference struct {
	// The FirewallPolicyArn of the FirewallPolicy resource.
	FirewallPolicyArn *string `field:"required" json:"firewallPolicyArn" yaml:"firewallPolicyArn"`
}

