package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customCodeSigningProperty := &CustomCodeSigningProperty{
//   	CertificateChain: &CodeSigningCertificateChainProperty{
//   		CertificateName: jsii.String("certificateName"),
//   		InlineDocument: jsii.String("inlineDocument"),
//   	},
//   	HashAlgorithm: jsii.String("hashAlgorithm"),
//   	Signature: &CodeSigningSignatureProperty{
//   		InlineDocument: jsii.String("inlineDocument"),
//   	},
//   	SignatureAlgorithm: jsii.String("signatureAlgorithm"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-customcodesigning.html
//
type CfnOTAUpdatePropsMixin_CustomCodeSigningProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-customcodesigning.html#cfn-iot-otaupdate-customcodesigning-certificatechain
	//
	CertificateChain interface{} `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// The hash algorithm used to code sign the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-customcodesigning.html#cfn-iot-otaupdate-customcodesigning-hashalgorithm
	//
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-customcodesigning.html#cfn-iot-otaupdate-customcodesigning-signature
	//
	Signature interface{} `field:"optional" json:"signature" yaml:"signature"`
	// The signature algorithm used to code sign the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-customcodesigning.html#cfn-iot-otaupdate-customcodesigning-signaturealgorithm
	//
	SignatureAlgorithm *string `field:"optional" json:"signatureAlgorithm" yaml:"signatureAlgorithm"`
}

