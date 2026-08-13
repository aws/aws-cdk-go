package awsec2


// Options to add a gateway endpoint to a VPC.
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
type GatewayVpcEndpointOptions struct {
	// The service to use for this gateway VPC endpoint.
	Service IGatewayVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Type of DNS records created for the VPC endpoint.
	// Default: not specified.
	//
	DnsRecordIpType VpcEndpointDnsRecordIpType `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// The IP address type for the endpoint.
	// Default: not specified.
	//
	IpAddressType VpcEndpointIpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Where to add endpoint routing.
	//
	// By default, this endpoint will be routable from all subnets in the VPC.
	// Specify a list of subnet selection objects here to be more specific.
	//
	// Example:
	//   var vpc Vpc
	//
	//
	//   vpc.addGatewayEndpoint(jsii.String("DynamoDbEndpoint"), &GatewayVpcEndpointOptions{
	//   	Service: ec2.GatewayVpcEndpointAwsService_DYNAMODB(),
	//   	// Add only to ISOLATED subnets
	//   	Subnets: []SubnetSelection{
	//   		&SubnetSelection{
	//   			SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
	//   		},
	//   	},
	//   })
	//
	// Default: - All subnets in the VPC.
	//
	Subnets *[]*SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
}

