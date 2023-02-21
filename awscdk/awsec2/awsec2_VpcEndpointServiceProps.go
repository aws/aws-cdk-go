package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties for a VpcEndpointService.
//
// Example:
//   var networkLoadBalancer1 networkLoadBalancer
//   var networkLoadBalancer2 networkLoadBalancer
//
//
//   ec2.NewVpcEndpointService(this, jsii.String("EndpointService"), &vpcEndpointServiceProps{
//   	vpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
//   		networkLoadBalancer1,
//   		networkLoadBalancer2,
//   	},
//   	acceptanceRequired: jsii.Boolean(true),
//   	allowedPrincipals: []arnPrincipal{
//   		iam.NewArnPrincipal(jsii.String("arn:aws:iam::123456789012:root")),
//   	},
//   })
//
// Experimental.
type VpcEndpointServiceProps struct {
	// One or more load balancers to host the VPC Endpoint Service.
	// Experimental.
	VpcEndpointServiceLoadBalancers *[]IVpcEndpointServiceLoadBalancer `field:"required" json:"vpcEndpointServiceLoadBalancers" yaml:"vpcEndpointServiceLoadBalancers"`
	// Whether requests from service consumers to connect to the service through an endpoint must be accepted.
	// Experimental.
	AcceptanceRequired *bool `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// IAM users, IAM roles, or AWS accounts to allow inbound connections from.
	//
	// These principals can connect to your service using VPC endpoints. Takes a
	// list of one or more ArnPrincipal.
	// Experimental.
	AllowedPrincipals *[]awsiam.ArnPrincipal `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
	// Name of the Vpc Endpoint Service.
	// Deprecated: This property is not used.
	VpcEndpointServiceName *string `field:"optional" json:"vpcEndpointServiceName" yaml:"vpcEndpointServiceName"`
	// IAM users, IAM roles, or AWS accounts to allow inbound connections from.
	//
	// These principals can connect to your service using VPC endpoints. Takes a
	// list of one or more ArnPrincipal.
	// Deprecated: use `allowedPrincipals`.
	WhitelistedPrincipals *[]awsiam.ArnPrincipal `field:"optional" json:"whitelistedPrincipals" yaml:"whitelistedPrincipals"`
}

