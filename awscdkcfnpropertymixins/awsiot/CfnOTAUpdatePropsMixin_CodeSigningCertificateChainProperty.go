package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeSigningCertificateChainProperty := &CodeSigningCertificateChainProperty{
//   	CertificateName: jsii.String("certificateName"),
//   	InlineDocument: jsii.String("inlineDocument"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigningcertificatechain.html
//
type CfnOTAUpdatePropsMixin_CodeSigningCertificateChainProperty struct {
	// The name of the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigningcertificatechain.html#cfn-iot-otaupdate-codesigningcertificatechain-certificatename
	//
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// A base64 encoded binary representation of the code signing certificate chain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-otaupdate-codesigningcertificatechain.html#cfn-iot-otaupdate-codesigningcertificatechain-inlinedocument
	//
	InlineDocument *string `field:"optional" json:"inlineDocument" yaml:"inlineDocument"`
}

