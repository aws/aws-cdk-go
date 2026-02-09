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
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type LoadBalancerContextResponse struct {
	// Type of IP address.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	IpAddressType LoadBalancerIpAddressType `field:"required" json:"ipAddressType" yaml:"ipAddressType"`
	// The ARN of the load balancer.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The hosted zone ID of the load balancer's name.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	LoadBalancerCanonicalHostedZoneId *string `field:"required" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// Load balancer's DNS name.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	LoadBalancerDnsName *string `field:"required" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Load balancer's security groups.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Load balancer's VPC.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

