package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Configures a custom domain endpoint for the Amazon OpenSearch Service domain.
//
// Example:
//   awscdk.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	customEndpoint: &customEndpointOptions{
//   		domainName: jsii.String("search.example.com"),
//   	},
//   })
//
type CustomEndpointOptions struct {
	// The custom domain name to assign.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The certificate to use.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The hosted zone in Route53 to create the CNAME record in.
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
}

