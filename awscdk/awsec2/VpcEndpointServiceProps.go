package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a VpcEndpointService.
//
// Example:
//   var networkLoadBalancer1 networkLoadBalancer
//   var networkLoadBalancer2 networkLoadBalancer
//
//
//   ec2.NewVpcEndpointService(this, jsii.String("EndpointService"), &VpcEndpointServiceProps{
//   	VpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
//   		networkLoadBalancer1,
//   		networkLoadBalancer2,
//   	},
//   	AcceptanceRequired: jsii.Boolean(true),
//   	AllowedPrincipals: []arnPrincipal{
//   		iam.NewArnPrincipal(jsii.String("arn:aws:iam::123456789012:root")),
//   	},
//   	ContributorInsights: jsii.Boolean(true),
//   })
//
type VpcEndpointServiceProps struct {
	// One or more load balancers to host the VPC Endpoint Service.
	VpcEndpointServiceLoadBalancers *[]IVpcEndpointServiceLoadBalancer `field:"required" json:"vpcEndpointServiceLoadBalancers" yaml:"vpcEndpointServiceLoadBalancers"`
	// Whether requests from service consumers to connect to the service through an endpoint must be accepted.
	// Default: true.
	//
	AcceptanceRequired *bool `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// IAM users, IAM roles, or AWS accounts to allow inbound connections from.
	//
	// These principals can connect to your service using VPC endpoints. Takes a
	// list of one or more ArnPrincipal.
	// Default: - no principals.
	//
	AllowedPrincipals *[]awsiam.ArnPrincipal `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
	// Indicates whether to enable the built-in Contributor Insights rules provided by AWS PrivateLink.
	// Default: false.
	//
	ContributorInsights *bool `field:"optional" json:"contributorInsights" yaml:"contributorInsights"`
}

