package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// properties used for creating the DomainName.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var certificate certificate
//
//   domainNameProps := &DomainNameProps{
//   	Certificate: certificate,
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	CertificateName: jsii.String("certificateName"),
//   	EndpointType: awscdk.Aws_apigatewayv2.EndpointType_EDGE,
//   	Mtls: &MTLSConfig{
//   		Bucket: bucket,
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	OwnershipCertificate: certificate,
//   	SecurityPolicy: awscdk.*Aws_apigatewayv2.SecurityPolicy_TLS_1_0,
//   }
//
type DomainNameProps struct {
	// The ACM certificate for this domain name.
	//
	// Certificate can be both ACM issued or imported.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The user-friendly name of the certificate that will be used by the endpoint for this domain name.
	// Default: - No friendly certificate name.
	//
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The type of endpoint for this DomainName.
	// Default: EndpointType.REGIONAL
	//
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// A public certificate issued by ACM to validate that you own a custom domain.
	//
	// This parameter is required
	// only when you configure mutual TLS authentication and you specify an ACM imported or private CA certificate
	// for `certificate`. The ownership certificate validates that you have permissions to use the domain name.
	// Default: - only required when configuring mTLS.
	//
	OwnershipCertificate awscertificatemanager.ICertificate `field:"optional" json:"ownershipCertificate" yaml:"ownershipCertificate"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// Default: SecurityPolicy.TLS_1_2
	//
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// The custom domain name.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The mutual TLS authentication configuration for a custom domain name.
	// Default: - mTLS is not configured.
	//
	Mtls *MTLSConfig `field:"optional" json:"mtls" yaml:"mtls"`
}

