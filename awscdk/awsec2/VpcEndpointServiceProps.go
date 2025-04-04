package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for a VpcEndpointService.
//
// Example:
//   var networkLoadBalancer networkLoadBalancer
//
//
//   ec2.NewVpcEndpointService(this, jsii.String("EndpointService"), &VpcEndpointServiceProps{
//   	VpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
//   		networkLoadBalancer,
//   	},
//   	// Support both IPv4 and IPv6 connections to the endpoint service
//   	SupportedIpAddressTypes: []ipAddressType{
//   		ec2.*ipAddressType_IPV4,
//   		ec2.*ipAddressType_IPV6,
//   	},
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
	// The Regions from which service consumers can access the service.
	// Default: - No Region restrictions.
	//
	AllowedRegions *[]*string `field:"optional" json:"allowedRegions" yaml:"allowedRegions"`
	// Indicates whether to enable the built-in Contributor Insights rules provided by AWS PrivateLink.
	// Default: false.
	//
	ContributorInsights *bool `field:"optional" json:"contributorInsights" yaml:"contributorInsights"`
	// Specify which IP address types are supported for VPC endpoint service.
	// Default: - No specific IP address types configured.
	//
	SupportedIpAddressTypes *[]IpAddressType `field:"optional" json:"supportedIpAddressTypes" yaml:"supportedIpAddressTypes"`
}

