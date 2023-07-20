package awselasticloadbalancingv2


// Specifies a subnet for a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetMappingProperty := &SubnetMappingProperty{
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	AllocationId: jsii.String("allocationId"),
//   	IPv6Address: jsii.String("iPv6Address"),
//   	PrivateIPv4Address: jsii.String("privateIPv4Address"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-subnetmapping.html
//
type CfnLoadBalancer_SubnetMappingProperty struct {
	// The ID of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-subnetmapping.html#cfn-elasticloadbalancingv2-loadbalancer-subnetmapping-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// [Network Load Balancers] The allocation ID of the Elastic IP address for an internet-facing load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-subnetmapping.html#cfn-elasticloadbalancingv2-loadbalancer-subnetmapping-allocationid
	//
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// [Network Load Balancers] The IPv6 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-subnetmapping.html#cfn-elasticloadbalancingv2-loadbalancer-subnetmapping-ipv6address
	//
	IPv6Address *string `field:"optional" json:"iPv6Address" yaml:"iPv6Address"`
	// [Network Load Balancers] The private IPv4 address for an internal load balancer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-subnetmapping.html#cfn-elasticloadbalancingv2-loadbalancer-subnetmapping-privateipv4address
	//
	PrivateIPv4Address *string `field:"optional" json:"privateIPv4Address" yaml:"privateIPv4Address"`
}

