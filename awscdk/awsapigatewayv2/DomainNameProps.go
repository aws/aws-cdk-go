package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// properties used for creating the DomainName.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   var bucket Bucket
//
//
//   certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
//   domainName := "example.com"
//
//   apigwv2.NewDomainName(this, jsii.String("DomainName"), &DomainNameProps{
//   	DomainName: jsii.String(DomainName),
//   	Certificate: acm.Certificate_FromCertificateArn(this, jsii.String("cert"), certArn),
//   	Mtls: &MTLSConfig{
//   		Bucket: *Bucket,
//   		Key: jsii.String("someca.pem"),
//   		Version: jsii.String("version"),
//   	},
//   })
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
	// The IP address types that can invoke the API.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-ip-address-type.html
	//
	// Default: undefined - AWS default is IPV4.
	//
	IpAddressType IpAddressType `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
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

