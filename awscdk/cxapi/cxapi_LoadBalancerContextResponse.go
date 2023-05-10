package cxapi


// Properties of a discovered load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerContextResponse := &LoadBalancerContextResponse{
//   	IpAddressType: awscdk.Cx_api.LoadBalancerIpAddressType_IPV4,
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	LoadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	LoadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// Experimental.
type LoadBalancerContextResponse struct {
	// Type of IP address.
	// Experimental.
	IpAddressType LoadBalancerIpAddressType `field:"required" json:"ipAddressType" yaml:"ipAddressType"`
	// The ARN of the load balancer.
	// Experimental.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The hosted zone ID of the load balancer's name.
	// Experimental.
	LoadBalancerCanonicalHostedZoneId *string `field:"required" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// Load balancer's DNS name.
	// Experimental.
	LoadBalancerDnsName *string `field:"required" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Load balancer's security groups.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Load balancer's VPC.
	// Experimental.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

