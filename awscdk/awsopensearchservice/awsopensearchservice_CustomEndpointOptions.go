package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

// Configures a custom domain endpoint for the Amazon OpenSearch Service domain.
//
// Example:
//   opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: opensearch.EngineVersion_OPENSEARCH_1_0(),
//   	CustomEndpoint: &CustomEndpointOptions{
//   		DomainName: jsii.String("search.example.com"),
//   	},
//   })
//
// Experimental.
type CustomEndpointOptions struct {
	// The custom domain name to assign.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The certificate to use.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The hosted zone in Route53 to create the CNAME record in.
	// Experimental.
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
}

