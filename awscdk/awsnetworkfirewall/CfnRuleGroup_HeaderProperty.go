package awsnetworkfirewall


// The 5-tuple criteria for AWS Network Firewall to use to inspect packet headers in stateful traffic flow inspection.
//
// Traffic flows that match the criteria are a match for the corresponding stateful rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerProperty := &HeaderProperty{
//   	Destination: jsii.String("destination"),
//   	DestinationPort: jsii.String("destinationPort"),
//   	Direction: jsii.String("direction"),
//   	Protocol: jsii.String("protocol"),
//   	Source: jsii.String("source"),
//   	SourcePort: jsii.String("sourcePort"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html
//
type CfnRuleGroup_HeaderProperty struct {
	// The destination IP address or address range to inspect for, in CIDR notation.
	//
	// To match with any address, specify `ANY` .
	//
	// Specify an IP address or a block of IP addresses in Classless Inter-Domain Routing (CIDR) notation. Network Firewall supports all address ranges for IPv4 and IPv6.
	//
	// Examples:
	//
	// - To configure Network Firewall to inspect for the IP address 192.0.2.44, specify `192.0.2.44/32` .
	// - To configure Network Firewall to inspect for IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24` .
	// - To configure Network Firewall to inspect for the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128` .
	// - To configure Network Firewall to inspect for IP addresses from 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64` .
	//
	// For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html#cfn-networkfirewall-rulegroup-header-destination
	//
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The destination port to inspect for.
	//
	// You can specify an individual port, for example `1994` and you can specify a port range, for example `1990:1994` . To match with any port, specify `ANY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html#cfn-networkfirewall-rulegroup-header-destinationport
	//
	DestinationPort *string `field:"required" json:"destinationPort" yaml:"destinationPort"`
	// The direction of traffic flow to inspect.
	//
	// If set to `ANY` , the inspection matches bidirectional traffic, both from the source to the destination and from the destination to the source. If set to `FORWARD` , the inspection only matches traffic going from the source to the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html#cfn-networkfirewall-rulegroup-header-direction
	//
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The protocol to inspect for.
	//
	// To specify all, you can use `IP` , because all traffic on AWS and on the internet is IP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html#cfn-networkfirewall-rulegroup-header-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The source IP address or address range to inspect for, in CIDR notation.
	//
	// To match with any address, specify `ANY` .
	//
	// Specify an IP address or a block of IP addresses in Classless Inter-Domain Routing (CIDR) notation. Network Firewall supports all address ranges for IPv4 and IPv6.
	//
	// Examples:
	//
	// - To configure Network Firewall to inspect for the IP address 192.0.2.44, specify `192.0.2.44/32` .
	// - To configure Network Firewall to inspect for IP addresses from 192.0.2.0 to 192.0.2.255, specify `192.0.2.0/24` .
	// - To configure Network Firewall to inspect for the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128` .
	// - To configure Network Firewall to inspect for IP addresses from 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64` .
	//
	// For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html#cfn-networkfirewall-rulegroup-header-source
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The source port to inspect for.
	//
	// You can specify an individual port, for example `1994` and you can specify a port range, for example `1990:1994` . To match with any port, specify `ANY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-header.html#cfn-networkfirewall-rulegroup-header-sourceport
	//
	SourcePort *string `field:"required" json:"sourcePort" yaml:"sourcePort"`
}

