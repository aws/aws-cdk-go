package awsec2


// Configuration for Vpc.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	IpAddresses: ec2.IpAddresses_Cidr(jsii.String("10.0.0.0/16")),
//   })
//
//   vpcConnector := apprunner.NewVpcConnector(this, jsii.String("VpcConnector"), &VpcConnectorProps{
//   	Vpc: Vpc,
//   	VpcSubnets: vpc.selectSubnets(&SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	}),
//   	VpcConnectorName: jsii.String("MyVpcConnector"),
//   })
//
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	VpcConnector: VpcConnector,
//   })
//
type VpcProps struct {
	// Availability zones this VPC spans.
	//
	// Specify this option only if you do not specify `maxAzs`.
	// Default: - a subset of AZs of the stack.
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The CIDR range to use for the VPC, e.g. '10.0.0.0/16'.
	//
	// Should be a minimum of /28 and maximum size of /16. The range will be
	// split across all subnets per Availability Zone.
	// Default: Vpc.DEFAULT_CIDR_RANGE
	//
	// Deprecated: Use ipAddresses instead.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// If set to false then disable the creation of the default internet gateway.
	// Default: true.
	//
	CreateInternetGateway *bool `field:"optional" json:"createInternetGateway" yaml:"createInternetGateway"`
	// The default tenancy of instances launched into the VPC.
	//
	// By setting this to dedicated tenancy, instances will be launched on
	// hardware dedicated to a single AWS customer, unless specifically specified
	// at instance launch time. Please note, not all instance types are usable
	// with Dedicated tenancy.
	// Default: DefaultInstanceTenancy.Default (shared) tenancy
	//
	DefaultInstanceTenancy DefaultInstanceTenancy `field:"optional" json:"defaultInstanceTenancy" yaml:"defaultInstanceTenancy"`
	// Indicates whether the instances launched in the VPC get public DNS hostnames.
	//
	// If this attribute is true, instances in the VPC get public DNS hostnames,
	// but only if the enableDnsSupport attribute is also set to true.
	// Default: true.
	//
	EnableDnsHostnames *bool `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	// Indicates whether the DNS resolution is supported for the VPC.
	//
	// If this attribute is false, the Amazon-provided DNS server in the VPC that
	// resolves public DNS hostnames to IP addresses is not enabled. If this
	// attribute is true, queries to the Amazon provided DNS server at the
	// 169.254.169.253 IP address, or the reserved IP address at the base of the
	// VPC IPv4 network range plus two will succeed.
	// Default: true.
	//
	EnableDnsSupport *bool `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	// Flow logs to add to this VPC.
	// Default: - No flow logs.
	//
	FlowLogs *map[string]*FlowLogOptions `field:"optional" json:"flowLogs" yaml:"flowLogs"`
	// Gateway endpoints to add to this VPC.
	// Default: - None.
	//
	GatewayEndpoints *map[string]*GatewayVpcEndpointOptions `field:"optional" json:"gatewayEndpoints" yaml:"gatewayEndpoints"`
	// The Provider to use to allocate IPv4 Space to your VPC.
	//
	// Options include static allocation or from a pool.
	//
	// Note this is specific to IPv4 addresses.
	// Default: ec2.IpAddresses.cidr
	//
	IpAddresses IIpAddresses `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// The protocol of the vpc.
	//
	// Options are IPv4 only or dual stack.
	// Default: IpProtocol.IPV4_ONLY
	//
	IpProtocol IpProtocol `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// The Provider to use to allocate IPv6 Space to your VPC.
	//
	// Options include amazon provided CIDR block.
	//
	// Note this is specific to IPv6 addresses.
	// Default: Ipv6Addresses.amazonProvided
	//
	Ipv6Addresses IIpv6Addresses `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Define the maximum number of AZs to use in this region.
	//
	// If the region has more AZs than you want to use (for example, because of
	// EIP limits), pick a lower number here. The AZs will be sorted and picked
	// from the start of the list.
	//
	// If you pick a higher number than the number of AZs in the region, all AZs
	// in the region will be selected. To use "all AZs" available to your
	// account, use a high number (such as 99).
	//
	// Be aware that environment-agnostic stacks will be created with access to
	// only 2 AZs, so to use more than 2 AZs, be sure to specify the account and
	// region on your stack.
	//
	// Specify this option only if you do not specify `availabilityZones`.
	// Default: 3.
	//
	MaxAzs *float64 `field:"optional" json:"maxAzs" yaml:"maxAzs"`
	// What type of NAT provider to use.
	//
	// Select between NAT gateways or NAT instances. NAT gateways
	// may not be available in all AWS regions.
	// Default: NatProvider.gateway()
	//
	NatGatewayProvider NatProvider `field:"optional" json:"natGatewayProvider" yaml:"natGatewayProvider"`
	// The number of NAT Gateways/Instances to create.
	//
	// The type of NAT gateway or instance will be determined by the
	// `natGatewayProvider` parameter.
	//
	// You can set this number lower than the number of Availability Zones in your
	// VPC in order to save on NAT cost. Be aware you may be charged for
	// cross-AZ data traffic instead.
	// Default: - One NAT gateway/instance per Availability Zone.
	//
	NatGateways *float64 `field:"optional" json:"natGateways" yaml:"natGateways"`
	// Configures the subnets which will have NAT Gateways/Instances.
	//
	// You can pick a specific group of subnets by specifying the group name;
	// the picked subnets must be public subnets.
	//
	// Only necessary if you have more than one public subnet group.
	// Default: - All public subnets.
	//
	NatGatewaySubnets *SubnetSelection `field:"optional" json:"natGatewaySubnets" yaml:"natGatewaySubnets"`
	// Define the number of AZs to reserve.
	//
	// When specified, the IP space is reserved for the azs but no actual
	// resources are provisioned.
	// Default: 0.
	//
	ReservedAzs *float64 `field:"optional" json:"reservedAzs" yaml:"reservedAzs"`
	// If set to true then the default inbound & outbound rules will be removed from the default security group.
	// Default: true if '@aws-cdk/aws-ec2:restrictDefaultSecurityGroup' is enabled, false otherwise.
	//
	RestrictDefaultSecurityGroup *bool `field:"optional" json:"restrictDefaultSecurityGroup" yaml:"restrictDefaultSecurityGroup"`
	// Configure the subnets to build for each AZ.
	//
	// Each entry in this list configures a Subnet Group; each group will contain a
	// subnet for each Availability Zone.
	//
	// For example, if you want 1 public subnet, 1 private subnet, and 1 isolated
	// subnet in each AZ provide the following:
	//
	// ```ts
	// new ec2.Vpc(this, 'VPC', {
	//   subnetConfiguration: [
	//      {
	//        cidrMask: 24,
	//        name: 'ingress',
	//        subnetType: ec2.SubnetType.PUBLIC,
	//      },
	//      {
	//        cidrMask: 24,
	//        name: 'application',
	//        subnetType: ec2.SubnetType.PRIVATE_WITH_EGRESS,
	//      },
	//      {
	//        cidrMask: 28,
	//        name: 'rds',
	//        subnetType: ec2.SubnetType.PRIVATE_ISOLATED,
	//      }
	//   ]
	// });
	// ```.
	// Default: - The VPC CIDR will be evenly divided between 1 public and 1
	// private subnet per AZ.
	//
	SubnetConfiguration *[]*SubnetConfiguration `field:"optional" json:"subnetConfiguration" yaml:"subnetConfiguration"`
	// The VPC name.
	//
	// Since the VPC resource doesn't support providing a physical name, the value provided here will be recorded in the `Name` tag.
	// Default: this.node.path
	//
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
	// VPN connections to this VPC.
	// Default: - No connections.
	//
	VpnConnections *map[string]*VpnConnectionOptions `field:"optional" json:"vpnConnections" yaml:"vpnConnections"`
	// Indicates whether a VPN gateway should be created and attached to this VPC.
	// Default: - true when vpnGatewayAsn or vpnConnections is specified.
	//
	VpnGateway *bool `field:"optional" json:"vpnGateway" yaml:"vpnGateway"`
	// The private Autonomous System Number (ASN) for the VPN gateway.
	// Default: - Amazon default ASN.
	//
	VpnGatewayAsn *float64 `field:"optional" json:"vpnGatewayAsn" yaml:"vpnGatewayAsn"`
	// Where to propagate VPN routes.
	// Default: - On the route tables associated with private subnets. If no
	// private subnets exists, isolated subnets are used. If no isolated subnets
	// exists, public subnets are used.
	//
	VpnRoutePropagation *[]*SubnetSelection `field:"optional" json:"vpnRoutePropagation" yaml:"vpnRoutePropagation"`
}

