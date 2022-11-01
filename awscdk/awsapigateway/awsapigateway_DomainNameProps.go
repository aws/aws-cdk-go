package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
)

// Example:
//   var acm interface{}
//
//
//   apigateway.NewDomainName(this, jsii.String("domain-name"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acm.certificate_FromCertificateArn(this, jsii.String("cert"), jsii.String("arn:aws:acm:us-east-1:1111111:certificate/11-3336f1-44483d-adc7-9cd375c5169d")),
//   	mtls: &mTLSConfig{
//   		bucket: s3.NewBucket(this, jsii.String("bucket")),
//   		key: jsii.String("truststore.pem"),
//   		version: jsii.String("version"),
//   	},
//   })
//
// Experimental.
type DomainNameProps struct {
	// The reference to an AWS-managed certificate for use by the edge-optimized endpoint for the domain name.
	//
	// For "EDGE" domain names, the certificate
	// needs to be in the US East (N. Virginia) region.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The custom domain name for your API.
	//
	// Uppercase letters are not supported.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	// Experimental.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The type of endpoint for this DomainName.
	// Experimental.
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The mutual TLS authentication configuration for a custom domain name.
	// Experimental.
	Mtls *MTLSConfig `field:"optional" json:"mtls" yaml:"mtls"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
	//
	// Experimental.
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// If specified, all requests to this domain will be mapped to the production deployment of this API.
	//
	// If you wish to map this domain to multiple APIs
	// with different base paths, don't specify this option and use
	// `addBasePathMapping`.
	// Experimental.
	Mapping IRestApi `field:"optional" json:"mapping" yaml:"mapping"`
}

