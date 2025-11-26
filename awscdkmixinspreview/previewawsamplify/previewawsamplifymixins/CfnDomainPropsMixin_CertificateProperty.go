package previewawsamplifymixins


// Describes the SSL/TLS certificate for the domain association.
//
// This can be your own custom certificate or the default certificate that Amplify provisions for you.
//
// If you are updating your domain to use a different certificate, `Certificate` points to the new certificate that is being created instead of the current active certificate. Otherwise, `Certificate` points to the current active certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   certificateProperty := &CertificateProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	CertificateType: jsii.String("certificateType"),
//   	CertificateVerificationDnsRecord: jsii.String("certificateVerificationDnsRecord"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html
//
type CfnDomainPropsMixin_CertificateProperty struct {
	// The Amazon resource name (ARN) for a custom certificate that you have already added to Certificate Manager in your AWS account .
	//
	// This field is required only when the certificate type is `CUSTOM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html#cfn-amplify-domain-certificate-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The type of SSL/TLS certificate that you want to use.
	//
	// Specify `AMPLIFY_MANAGED` to use the default certificate that Amplify provisions for you.
	//
	// Specify `CUSTOM` to use your own certificate that you have already added to Certificate Manager in your AWS account . Make sure you request (or import) the certificate in the US East (N. Virginia) Region (us-east-1). For more information about using ACM, see [Importing certificates into Certificate Manager](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *ACM User guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html#cfn-amplify-domain-certificate-certificatetype
	//
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
	// The DNS record for certificate verification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificate.html#cfn-amplify-domain-certificate-certificateverificationdnsrecord
	//
	CertificateVerificationDnsRecord *string `field:"optional" json:"certificateVerificationDnsRecord" yaml:"certificateVerificationDnsRecord"`
}

