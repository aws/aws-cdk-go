package previewawsec2mixins


// Public hostname type options.
//
// For more information, see [EC2 instance hostnames, DNS names, and domains](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publicIpDnsNameOptionsProperty := &PublicIpDnsNameOptionsProperty{
//   	DnsHostnameType: jsii.String("dnsHostnameType"),
//   	PublicDualStackDnsName: jsii.String("publicDualStackDnsName"),
//   	PublicIpv4DnsName: jsii.String("publicIpv4DnsName"),
//   	PublicIpv6DnsName: jsii.String("publicIpv6DnsName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-publicipdnsnameoptions.html
//
type CfnNetworkInterfacePropsMixin_PublicIpDnsNameOptionsProperty struct {
	// The public hostname type.
	//
	// For more information, see [EC2 instance hostnames, DNS names, and domains](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-publicipdnsnameoptions.html#cfn-ec2-networkinterface-publicipdnsnameoptions-dnshostnametype
	//
	DnsHostnameType *string `field:"optional" json:"dnsHostnameType" yaml:"dnsHostnameType"`
	// A dual-stack public hostname for a network interface.
	//
	// Requests from within the VPC resolve to both the private IPv4 address and the IPv6 Global Unicast Address of the network interface. Requests from the internet resolve to both the public IPv4 and the IPv6 GUA address of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-publicipdnsnameoptions.html#cfn-ec2-networkinterface-publicipdnsnameoptions-publicdualstackdnsname
	//
	PublicDualStackDnsName *string `field:"optional" json:"publicDualStackDnsName" yaml:"publicDualStackDnsName"`
	// An IPv4-enabled public hostname for a network interface.
	//
	// Requests from within the VPC resolve to the private primary IPv4 address of the network interface. Requests from the internet resolve to the public IPv4 address of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-publicipdnsnameoptions.html#cfn-ec2-networkinterface-publicipdnsnameoptions-publicipv4dnsname
	//
	PublicIpv4DnsName *string `field:"optional" json:"publicIpv4DnsName" yaml:"publicIpv4DnsName"`
	// An IPv6-enabled public hostname for a network interface.
	//
	// Requests from within the VPC or from the internet resolve to the IPv6 GUA of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-publicipdnsnameoptions.html#cfn-ec2-networkinterface-publicipdnsnameoptions-publicipv6dnsname
	//
	PublicIpv6DnsName *string `field:"optional" json:"publicIpv6DnsName" yaml:"publicIpv6DnsName"`
}

