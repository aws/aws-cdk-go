package previewawsapigatewayv2mixins


// The `DomainNameConfiguration` property type specifies the configuration for an API's domain name.
//
// `DomainNameConfiguration` is a property of the [AWS::ApiGatewayV2::DomainName](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-domainname.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   domainNameConfigurationProperty := &DomainNameConfigurationProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	CertificateName: jsii.String("certificateName"),
//   	EndpointType: jsii.String("endpointType"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	OwnershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	SecurityPolicy: jsii.String("securityPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html
//
type CfnDomainNamePropsMixin_DomainNameConfigurationProperty struct {
	// An AWS -managed certificate that will be used by the edge-optimized endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The user-friendly name of the certificate that will be used by the edge-optimized endpoint for this domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-certificatename
	//
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// The endpoint type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-endpointtype
	//
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The IP address types that can invoke the domain name.
	//
	// Use `ipv4` to allow only IPv4 addresses to invoke your domain name, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your domain name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The Amazon resource name (ARN) for the public certificate issued by Certificate Manager .
	//
	// This ARN is used to validate custom domain ownership. It's required only if you configure mutual TLS and use either an ACM-imported or a private CA certificate ARN as the regionalCertificateArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-ownershipverificationcertificatearn
	//
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// The Transport Layer Security (TLS) version of the security policy for this domain name.
	//
	// The valid values are `TLS_1_0` and `TLS_1_2` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-domainname-domainnameconfiguration.html#cfn-apigatewayv2-domainname-domainnameconfiguration-securitypolicy
	//
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

