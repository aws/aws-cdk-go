package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscertificatemanager"
)

// Configures a custom domain endpoint for the Amazon OpenSearch Service domain.
//
// Example:
//   awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	CustomEndpoint: &CustomEndpointOptions{
//   		DomainName: jsii.String("search.example.com"),
//   	},
//   })
//
type CustomEndpointOptions struct {
	// The custom domain name to assign.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The certificate to use.
	// Default: - create a new one.
	//
	Certificate interfacesawscertificatemanager.ICertificateRef `field:"optional" json:"certificate" yaml:"certificate"`
	// The hosted zone in Route53 to create the CNAME record in.
	// Default: - do not create a CNAME.
	//
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
}

