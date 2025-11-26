package previewawsacmpcamixins


// Properties for CfnCertificateAuthorityActivationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateAuthorityActivationMixinProps := &CfnCertificateAuthorityActivationMixinProps{
//   	Certificate: jsii.String("certificate"),
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	CertificateChain: jsii.String("certificateChain"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html
//
type CfnCertificateAuthorityActivationMixinProps struct {
	// The Base64 PEM-encoded certificate authority certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The Amazon Resource Name (ARN) of your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// The Base64 PEM-encoded certificate chain that chains up to the root CA certificate that you used to sign your private CA certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificatechain
	//
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Status of your private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

