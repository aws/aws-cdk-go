package previewawsnetworkfirewallmixins


// Settings that define the Secure Sockets Layer/Transport Layer Security (SSL/TLS) traffic that Network Firewall should decrypt for inspection by the stateful rule engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverCertificateScopeProperty := &ServerCertificateScopeProperty{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.html
//
type CfnTLSInspectionConfigurationPropsMixin_ServerCertificateScopeProperty struct {
	// The destination ports to decrypt for inspection, in Transmission Control Protocol (TCP) format.
	//
	// If not specified, this matches with any destination port.
	//
	// You can specify individual ports, for example `1994` , and you can specify port ranges, such as `1990:1994` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificatescope-destinationports
	//
	DestinationPorts interface{} `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// The destination IP addresses and address ranges to decrypt for inspection, in CIDR notation.
	//
	// If not specified, this
	// matches with any destination address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificatescope-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The protocols to inspect for, specified using the assigned internet protocol number (IANA) for each protocol.
	//
	// If not specified, this matches with any protocol.
	//
	// Network Firewall currently supports only TCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificatescope-protocols
	//
	Protocols interface{} `field:"optional" json:"protocols" yaml:"protocols"`
	// The source ports to decrypt for inspection, in Transmission Control Protocol (TCP) format.
	//
	// If not specified, this matches with any source port.
	//
	// You can specify individual ports, for example `1994` , and you can specify port ranges, such as `1990:1994` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificatescope-sourceports
	//
	SourcePorts interface{} `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
	// The source IP addresses and address ranges to decrypt for inspection, in CIDR notation.
	//
	// If not specified, this
	// matches with any source address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-tlsinspectionconfiguration-servercertificatescope.html#cfn-networkfirewall-tlsinspectionconfiguration-servercertificatescope-sources
	//
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

