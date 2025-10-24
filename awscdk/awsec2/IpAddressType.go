package awsec2


// IP address types supported for VPC endpoint service.
//
// Example:
//   var networkLoadBalancer NetworkLoadBalancer
//
//
//   ec2.NewVpcEndpointService(this, jsii.String("EndpointService"), &VpcEndpointServiceProps{
//   	VpcEndpointServiceLoadBalancers: []IVpcEndpointServiceLoadBalancer{
//   		networkLoadBalancer,
//   	},
//   	// Support both IPv4 and IPv6 connections to the endpoint service
//   	SupportedIpAddressTypes: []IpAddressType{
//   		ec2.IpAddressType_IPV4,
//   		ec2.IpAddressType_IPV6,
//   	},
//   })
//
type IpAddressType string

const (
	// ipv4 address type.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// ipv6 address type.
	IpAddressType_IPV6 IpAddressType = "IPV6"
)

