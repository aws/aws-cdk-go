package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// properties used for creating the DomainName.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws-samples/dummy/awscdklib/awsapigatewayv2integrations"
//
//   var handler function
//
//
//   certArn := "arn:aws:acm:us-east-1:111111111111:certificate"
//   domainName := "example.com"
//
//   dn := apigwv2.NewDomainName(this, jsii.String("DN"), &DomainNameProps{
//   	DomainName: domainName,
//   	Certificate: acm.Certificate_FromCertificateArn(this, jsii.String("cert"), certArn),
//   })
//   api := apigwv2.NewHttpApi(this, jsii.String("HttpProxyProdApi"), &HttpApiProps{
//   	DefaultIntegration: awscdklibawsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("DefaultIntegration"), handler),
//   	// https://${dn.domainName}/foo goes to prodApi $default stage
//   	DefaultDomainMapping: &DomainMappingOptions{
//   		DomainName: dn,
//   		MappingKey: jsii.String("foo"),
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

