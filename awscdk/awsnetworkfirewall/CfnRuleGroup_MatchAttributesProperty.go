package awsnetworkfirewall


// Criteria for Network Firewall to use to inspect an individual packet in stateless rule inspection.
//
// Each match attributes set can include one or more items such as IP address, CIDR range, port number, protocol, and TCP flags.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchAttributesProperty := &MatchAttributesProperty{
//   	DestinationPorts: []interface{}{
//   		&PortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Destinations: []interface{}{
//   		&AddressProperty{
//   			AddressDefinition: jsii.String("addressDefinition"),
//   		},
//   	},
//   	Protocols: []interface{}{
//   		jsii.Number(123),
//   	},
//   	SourcePorts: []interface{}{
//   		&PortRangeProperty{
//   			FromPort: jsii.Number(123),
//   			ToPort: jsii.Number(123),
//   		},
//   	},
//   	Sources: []interface{}{
//   		&AddressProperty{
//   			AddressDefinition: jsii.String("addressDefinition"),
//   		},
//   	},
//   	TcpFlags: []interface{}{
//   		&TCPFlagFieldProperty{
//   			Flags: []*string{
//   				jsii.String("flags"),
//   			},
//
//   			// the properties below are optional
//   			Masks: []*string{
//   				jsii.String("masks"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleGroup_MatchAttributesProperty struct {
	// The destination ports to inspect for.
	//
	// If not specified, this matches with any destination port. This setting is only used for protocols 6 (TCP) and 17 (UDP).
	//
	// You can specify individual ports, for example `1994` and you can specify port ranges, for example `1990:1994` .
	DestinationPorts interface{} `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// The destination IP addresses and address ranges to inspect for, in CIDR notation.
	//
	// If not specified, this matches with any destination address.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The protocols to inspect for, specified using each protocol's assigned internet protocol number (IANA).
	//
	// If not specified, this matches with any protocol.
	Protocols interface{} `field:"optional" json:"protocols" yaml:"protocols"`
	// The source ports to inspect for.
	//
	// If not specified, this matches with any source port. This setting is only used for protocols 6 (TCP) and 17 (UDP).
	//
	// You can specify individual ports, for example `1994` and you can specify port ranges, for example `1990:1994` .
	SourcePorts interface{} `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
	// The source IP addresses and address ranges to inspect for, in CIDR notation.
	//
	// If not specified, this matches with any source address.
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
	// The TCP flags and masks to inspect for.
	//
	// If not specified, this matches with any settings. This setting is only used for protocol 6 (TCP).
	TcpFlags interface{} `field:"optional" json:"tcpFlags" yaml:"tcpFlags"`
}

