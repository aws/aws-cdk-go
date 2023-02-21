package awselasticloadbalancingv2


// Specifies a subnet for a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetMappingProperty := &subnetMappingProperty{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	allocationId: jsii.String("allocationId"),
//   	iPv6Address: jsii.String("iPv6Address"),
//   	privateIPv4Address: jsii.String("privateIPv4Address"),
//   }
//
type CfnLoadBalancer_SubnetMappingProperty struct {
	// The ID of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// [Network Load Balancers] The IPv6 address.
	IPv6Address *string `field:"optional" json:"iPv6Address" yaml:"iPv6Address"`
	// [Network Load Balancers] The private IPv4 address for an internal load balancer.
	PrivateIPv4Address *string `field:"optional" json:"privateIPv4Address" yaml:"privateIPv4Address"`
}

