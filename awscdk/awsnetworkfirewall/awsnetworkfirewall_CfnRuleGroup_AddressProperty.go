package awsnetworkfirewall


// A single IP address specification.
//
// This is used in the `RuleGroup.MatchAttributes` source and destination specifications.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addressProperty := &addressProperty{
//   	addressDefinition: jsii.String("addressDefinition"),
//   }
//
type CfnRuleGroup_AddressProperty struct {
	// Specify an IP address or a block of IP addresses in Classless Inter-Domain Routing (CIDR) notation.
	//
	// Network Firewall supports all address ranges for IPv4.
	//
	// Examples:
	//
	// - To configure Network Firewall to inspect for the IP address 192.0.2.44, specify `192.0.2.44/32` .
	// - To configure Network Firewall to inspect for IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24` .
	//
	// For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
	AddressDefinition *string `field:"required" json:"addressDefinition" yaml:"addressDefinition"`
}

