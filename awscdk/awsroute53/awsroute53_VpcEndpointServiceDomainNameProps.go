package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties to configure a VPC Endpoint Service domain name.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var zone hostedZone
//   var vpces vpcEndpointService
//
//
//   awscdk.NewVpcEndpointServiceDomainName(this, jsii.String("EndpointDomain"), &VpcEndpointServiceDomainNameProps{
//   	EndpointService: vpces,
//   	DomainName: jsii.String("my-stuff.aws-cdk.dev"),
//   	PublicHostedZone: zone,
//   })
//
// Experimental.
type VpcEndpointServiceDomainNameProps struct {
	// The domain name to use.
	//
	// This domain name must be owned by this account (registered through Route53),
	// or delegated to this account. Domain ownership will be verified by AWS before
	// private DNS can be used.
	// See: https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html
	//
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The VPC Endpoint Service to configure Private DNS for.
	// Experimental.
	EndpointService awsec2.IVpcEndpointService `field:"required" json:"endpointService" yaml:"endpointService"`
	// The public hosted zone to use for the domain.
	// Experimental.
	PublicHostedZone IPublicHostedZone `field:"required" json:"publicHostedZone" yaml:"publicHostedZone"`
}

