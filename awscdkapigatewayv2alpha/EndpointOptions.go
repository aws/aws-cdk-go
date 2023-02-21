// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// properties for creating a domain name endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificate certificate
//
//   endpointOptions := &EndpointOptions{
//   	Certificate: certificate,
//
//   	// the properties below are optional
//   	CertificateName: jsii.String("certificateName"),
//   	EndpointType: apigatewayv2_alpha.EndpointType_EDGE,
//   	OwnershipCertificate: certificate,
//   	SecurityPolicy: apigatewayv2_alpha.SecurityPolicy_TLS_1_0,
//   }
//
// Experimental.
type EndpointOptions struct {
	// The ACM certificate for this domain name.
	//
	// Certificate can be both ACM issued or imported.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The user-friendly name of the certificate that will be used by the endpoint for this domain name.
	// Experimental.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The type of endpoint for this DomainName.
	// Experimental.
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// A public certificate issued by ACM to validate that you own a custom domain.
	//
	// This parameter is required
	// only when you configure mutual TLS authentication and you specify an ACM imported or private CA certificate
	// for `certificate`. The ownership certificate validates that you have permissions to use the domain name.
	// Experimental.
	OwnershipCertificate awscertificatemanager.ICertificate `field:"optional" json:"ownershipCertificate" yaml:"ownershipCertificate"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// Experimental.
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

