package awsec2


// Enums for all Dns Record IP Address types.
//
// Example:
//   // Interface endpoint with IPv6
//   vpc.addInterfaceEndpoint(jsii.String("ExampleEndpoint"), &InterfaceVpcEndpointOptions{
//   	Service: ec2.InterfaceVpcEndpointAwsService_ECR(),
//   	IpAddressType: ec2.VpcEndpointIpAddressType_IPV6,
//   	DnsRecordIpType: ec2.VpcEndpointDnsRecordIpType_IPV6,
//   })
//
//   // Gateway endpoint with dualstack
//   vpc.addGatewayEndpoint(jsii.String("S3DualstackEndpoint"), &GatewayVpcEndpointOptions{
//   	Service: ec2.GatewayVpcEndpointAwsService_S3(),
//   	IpAddressType: ec2.VpcEndpointIpAddressType_DUALSTACK,
//   	DnsRecordIpType: ec2.VpcEndpointDnsRecordIpType_DUALSTACK,
//   })
//
type VpcEndpointDnsRecordIpType string

const (
	// Create A records for the private, Regional, and zonal DNS names.
	//
	// The IP address type must be IPv4 or Dualstack.
	VpcEndpointDnsRecordIpType_IPV4 VpcEndpointDnsRecordIpType = "IPV4"
	// Create AAAA records for the private, Regional, and zonal DNS names.
	//
	// The IP address type must be IPv6 or Dualstack.
	VpcEndpointDnsRecordIpType_IPV6 VpcEndpointDnsRecordIpType = "IPV6"
	// Create A and AAAA records for the private, Regional, and zonal DNS names.
	//
	// The IP address type must be Dualstack.
	VpcEndpointDnsRecordIpType_DUALSTACK VpcEndpointDnsRecordIpType = "DUALSTACK"
	// Create A records for the private, Regional, and zonal DNS names and AAAA records for the Regional and zonal DNS names.
	//
	// The IP address type must be Dualstack.
	VpcEndpointDnsRecordIpType_SERVICE_DEFINED VpcEndpointDnsRecordIpType = "SERVICE_DEFINED"
)

