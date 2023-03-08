package awsapigatewayv2


// The `DomainNameConfiguration` property type specifies the configuration for a an API's domain name.
//
// `DomainNameConfiguration` is a property of the [AWS::ApiGatewayV2::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainNameConfigurationProperty := &domainNameConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	certificateName: jsii.String("certificateName"),
//   	endpointType: jsii.String("endpointType"),
//   	ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	securityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnDomainName_DomainNameConfigurationProperty struct {
	// An AWS -managed certificate that will be used by the edge-optimized endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The endpoint type.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The Amazon resource name (ARN) for the public certificate issued by AWS Certificate Manager .
	//
	// This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// The Transport Layer Security (TLS) version of the security policy for this domain name.
	//
	// The valid values are `TLS_1_0` and `TLS_1_2` .
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

