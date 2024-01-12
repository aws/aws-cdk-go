package awsec2


// The types of IP addresses provisioned in the VPC.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("TheVPC"), &VpcProps{
//   	IpProtocol: ec2.IpProtocol_DUAL_STACK,
//
//   	SubnetConfiguration: []subnetConfiguration{
//   		&subnetConfiguration{
//   			// general properties
//   			Name: jsii.String("Public"),
//   			SubnetType: ec2.SubnetType_PUBLIC,
//   			Reserved: jsii.Boolean(false),
//
//   			// IPv4 specific properties
//   			MapPublicIpOnLaunch: jsii.Boolean(true),
//   			CidrMask: jsii.Number(24),
//
//   			// new IPv6 specific property
//   			Ipv6AssignAddressOnCreation: jsii.Boolean(true),
//   		},
//   	},
//   })
//
type IpProtocol string

const (
	// The vpc will be configured with only IPv4 addresses.
	//
	// This is the default protocol if unset.
	IpProtocol_IPV4_ONLY IpProtocol = "IPV4_ONLY"
	// The vpc will have both IPv4 and IPv6 addresses.
	//
	// Unless specified, public IPv4 addresses will not be auto assigned,
	// an egress only internet gateway (EIGW) will be created and configured,
	// and NATs and internet gateways (IGW) will be configured with IPv6 addresses.
	IpProtocol_DUAL_STACK IpProtocol = "DUAL_STACK"
)

