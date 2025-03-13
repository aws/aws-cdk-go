package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Configures a custom domain endpoint for the ES domain.
//
// Example:
//   es.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: es.ElasticsearchVersion_V7_7(),
//   	CustomEndpoint: &CustomEndpointOptions{
//   		DomainName: jsii.String("search.example.com"),
//   	},
//   })
//
// Deprecated: use opensearchservice module instead.
type CustomEndpointOptions struct {
	// The custom domain name to assign.
	// Deprecated: use opensearchservice module instead.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The certificate to use.
	// Default: - create a new one.
	//
	// Deprecated: use opensearchservice module instead.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The hosted zone in Route53 to create the CNAME record in.
	// Default: - do not create a CNAME.
	//
	// Deprecated: use opensearchservice module instead.
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
}

