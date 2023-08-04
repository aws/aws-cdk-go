package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to reference an existing load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   applicationLoadBalancerAttributes := &ApplicationLoadBalancerAttributes{
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	SecurityGroupId: jsii.String("securityGroupId"),
//
//   	// the properties below are optional
//   	LoadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	LoadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	SecurityGroupAllowsAllOutbound: jsii.Boolean(false),
//   	Vpc: vpc,
//   }
//
type ApplicationLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// ID of the load balancer's security group.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The canonical hosted zone ID of this load balancer.
	// Default: - When not provided, LB cannot be used as Route53 Alias target.
	//
	LoadBalancerCanonicalHostedZoneId *string `field:"optional" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	// Default: - When not provided, LB cannot be used as Route53 Alias target.
	//
	LoadBalancerDnsName *string `field:"optional" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Whether the security group allows all outbound traffic or not.
	//
	// Unless set to `false`, no egress rules will be added to the security group.
	// Default: true.
	//
	SecurityGroupAllowsAllOutbound *bool `field:"optional" json:"securityGroupAllowsAllOutbound" yaml:"securityGroupAllowsAllOutbound"`
	// The VPC this load balancer has been created in, if available.
	// Default: - If the Load Balancer was imported and a VPC was not specified,
	// the VPC is not available.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

