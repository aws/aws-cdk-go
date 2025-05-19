package awsec2


// IP address type for the endpoint.
//
// Example:
//   vpc.addInterfaceEndpoint(jsii.String("ExampleEndpoint"), &InterfaceVpcEndpointOptions{
//   	Service: ec2.InterfaceVpcEndpointAwsService_ECR(),
//   	IpAddressType: ec2.VpcEndpointIpAddressType_IPV6,
//   	DnsRecordIpType: ec2.VpcEndpointDnsRecordIpType_IPV6,
//   })
//
type VpcEndpointIpAddressType string

const (
	// Assign IPv4 addresses to the endpoint network interfaces.
	//
	// This option is supported only if all selected subnets have IPv4 address ranges
	// and the endpoint service accepts IPv4 requests.
	VpcEndpointIpAddressType_IPV4 VpcEndpointIpAddressType = "IPV4"
	// Assign IPv6 addresses to the endpoint network interfaces.
	//
	// This option is supported only if all selected subnets are IPv6 only subnets
	// and the endpoint service accepts IPv6 requests.
	VpcEndpointIpAddressType_IPV6 VpcEndpointIpAddressType = "IPV6"
	// Assign both IPv4 and IPv6 addresses to the endpoint network interfaces.
	//
	// This option is supported only if all selected subnets have both IPv4 and IPv6
	// address ranges and the endpoint service accepts both IPv4 and IPv6 requests.
	VpcEndpointIpAddressType_DUALSTACK VpcEndpointIpAddressType = "DUALSTACK"
)

