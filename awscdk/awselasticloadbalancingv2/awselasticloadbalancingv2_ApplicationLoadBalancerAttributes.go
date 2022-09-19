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
//   applicationLoadBalancerAttributes := &applicationLoadBalancerAttributes{
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	securityGroupId: jsii.String("securityGroupId"),
//
//   	// the properties below are optional
//   	loadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	loadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	securityGroupAllowsAllOutbound: jsii.Boolean(false),
//   	vpc: vpc,
//   }
//
type ApplicationLoadBalancerAttributes struct {
	// ARN of the load balancer.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// ID of the load balancer's security group.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The canonical hosted zone ID of this load balancer.
	LoadBalancerCanonicalHostedZoneId *string `field:"optional" json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// The DNS name of this load balancer.
	LoadBalancerDnsName *string `field:"optional" json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Whether the security group allows all outbound traffic or not.
	//
	// Unless set to `false`, no egress rules will be added to the security group.
	SecurityGroupAllowsAllOutbound *bool `field:"optional" json:"securityGroupAllowsAllOutbound" yaml:"securityGroupAllowsAllOutbound"`
	// The VPC this load balancer has been created in, if available.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

