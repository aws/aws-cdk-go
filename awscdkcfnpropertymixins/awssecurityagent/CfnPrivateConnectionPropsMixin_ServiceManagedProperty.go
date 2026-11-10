package awssecurityagent


// The configuration for a service-managed private connection, where the service manages the resource gateway lifecycle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   serviceManagedProperty := &ServiceManagedProperty{
//   	DnsResolution: jsii.String("dnsResolution"),
//   	HostAddress: jsii.String("hostAddress"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Ipv4AddressesPerEni: jsii.Number(123),
//   	PortRanges: []*string{
//   		jsii.String("portRanges"),
//   	},
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html
//
type CfnPrivateConnectionPropsMixin_ServiceManagedProperty struct {
	// The DNS resolution mode for the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-dnsresolution
	//
	DnsResolution *string `field:"optional" json:"dnsResolution" yaml:"dnsResolution"`
	// The IP address or DNS name of the target resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-hostaddress
	//
	HostAddress *string `field:"optional" json:"hostAddress" yaml:"hostAddress"`
	// The IP address type of the service-managed resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The number of IPv4 addresses in each elastic network interface for the service-managed resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-ipv4addressespereni
	//
	Ipv4AddressesPerEni *float64 `field:"optional" json:"ipv4AddressesPerEni" yaml:"ipv4AddressesPerEni"`
	// The TCP port ranges that a consumer can use to access the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-portranges
	//
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// The security groups to attach to the service-managed resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnets that the service-managed resource gateway spans.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The VPC to create the service-managed resource gateway in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-servicemanaged.html#cfn-securityagent-privateconnection-servicemanaged-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

