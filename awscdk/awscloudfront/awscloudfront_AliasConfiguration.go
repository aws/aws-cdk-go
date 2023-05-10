package awscloudfront


// Configuration for custom domain names.
//
// CloudFront can use a custom domain that you provide instead of a
// "cloudfront.net" domain. To use this feature you must provide the list of
// additional domains, and the ACM Certificate that CloudFront should use for
// these additional domains.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aliasConfiguration := &AliasConfiguration{
//   	AcmCertRef: jsii.String("acmCertRef"),
//   	Names: []*string{
//   		jsii.String("names"),
//   	},
//
//   	// the properties below are optional
//   	SecurityPolicy: awscdk.Aws_cloudfront.SecurityPolicyProtocol_SSL_V3,
//   	SslMethod: awscdk.*Aws_cloudfront.SSLMethod_SNI,
//   }
//
// Deprecated: see {@link CloudFrontWebDistributionProps#viewerCertificate} with {@link ViewerCertificate#acmCertificate}.
type AliasConfiguration struct {
	// ARN of an AWS Certificate Manager (ACM) certificate.
	// Deprecated: see {@link CloudFrontWebDistributionProps#viewerCertificate} with {@link ViewerCertificate#acmCertificate}.
	AcmCertRef *string `field:"required" json:"acmCertRef" yaml:"acmCertRef"`
	// Domain names on the certificate.
	//
	// Both main domain name and Subject Alternative Names.
	// Deprecated: see {@link CloudFrontWebDistributionProps#viewerCertificate} with {@link ViewerCertificate#acmCertificate}.
	Names *[]*string `field:"required" json:"names" yaml:"names"`
	// The minimum version of the SSL protocol that you want CloudFront to use for HTTPS connections.
	//
	// CloudFront serves your objects only to browsers or devices that support at
	// least the SSL version that you specify.
	// Deprecated: see {@link CloudFrontWebDistributionProps#viewerCertificate} with {@link ViewerCertificate#acmCertificate}.
	SecurityPolicy SecurityPolicyProtocol `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// How CloudFront should serve HTTPS requests.
	//
	// See the notes on SSLMethod if you wish to use other SSL termination types.
	// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ViewerCertificate.html
	//
	// Deprecated: see {@link CloudFrontWebDistributionProps#viewerCertificate} with {@link ViewerCertificate#acmCertificate}.
	SslMethod SSLMethod `field:"optional" json:"sslMethod" yaml:"sslMethod"`
}

