package awsdevopsagent


// Configuration for a service-managed Private Connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceManagedModeProperty := &ServiceManagedModeProperty{
//   	HostAddress: jsii.String("hostAddress"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html
//
type CfnPrivateConnection_ServiceManagedModeProperty struct {
	// IP address or DNS name of the target resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-hostaddress
	//
	HostAddress *string `field:"required" json:"hostAddress" yaml:"hostAddress"`
	// VPC to create the service-managed Resource Gateway in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// IP address type of the service-managed Resource Gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Number of IPv4 addresses in each ENI for the service-managed Resource Gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-ipv4addressespereni
	//
	Ipv4AddressesPerEni *float64 `field:"optional" json:"ipv4AddressesPerEni" yaml:"ipv4AddressesPerEni"`
	// TCP port ranges that a consumer can use to access the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-portranges
	//
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// Security groups to attach to the service-managed Resource Gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Subnets that the service-managed Resource Gateway will span.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-servicemanagedmode.html#cfn-devopsagent-privateconnection-servicemanagedmode-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

