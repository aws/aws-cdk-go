package interfacesawsnetworkfirewall


// A reference to a Firewall resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firewallReference := &FirewallReference{
//   	FirewallArn: jsii.String("firewallArn"),
//   }
//
type FirewallReference struct {
	// The FirewallArn of the Firewall resource.
	FirewallArn *string `field:"required" json:"firewallArn" yaml:"firewallArn"`
}

