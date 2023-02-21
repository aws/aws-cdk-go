package cxapi


// Properties of a discovered load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerContextResponse := &loadBalancerContextResponse{
//   	ipAddressType: awscdk.Cx_api.loadBalancerIpAddressType_IPV4,
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	loadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
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

