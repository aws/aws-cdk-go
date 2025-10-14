package awsamplify


// The type of SSL/TLS certificate to use for your custom domain.
//
// If a certificate type isn't specified, Amplify uses the default `AMPLIFY_MANAGED` certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateSettingsProperty := &CertificateSettingsProperty{
//   	CertificateType: jsii.String("certificateType"),
//   	CustomCertificateArn: jsii.String("customCertificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificatesettings.html
//
type CfnDomain_CertificateSettingsProperty struct {
	// The certificate type.
	//
	// Specify `AMPLIFY_MANAGED` to use the default certificate that Amplify provisions for you.
	//
	// Specify `CUSTOM` to use your own certificate that you have already added to Certificate Manager in your AWS account . Make sure you request (or import) the certificate in the US East (N. Virginia) Region (us-east-1). For more information about using ACM, see [Importing certificates into Certificate Manager](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html) in the *ACM User guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificatesettings.html#cfn-amplify-domain-certificatesettings-certificatetype
	//
	CertificateType *string `field:"optional" json:"certificateType" yaml:"certificateType"`
	// The Amazon resource name (ARN) for the custom certificate that you have already added to Certificate Manager in your AWS account .
	//
	// This field is required only when the certificate type is `CUSTOM` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplify-domain-certificatesettings.html#cfn-amplify-domain-certificatesettings-customcertificatearn
	//
	CustomCertificateArn *string `field:"optional" json:"customCertificateArn" yaml:"customCertificateArn"`
}

