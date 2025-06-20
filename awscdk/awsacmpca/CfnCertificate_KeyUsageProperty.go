package awsacmpca


// Defines one or more purposes for which the key contained in the certificate can be used.
//
// Default value for each option is false.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyUsageProperty := &KeyUsageProperty{
//   	CrlSign: jsii.Boolean(false),
//   	DataEncipherment: jsii.Boolean(false),
//   	DecipherOnly: jsii.Boolean(false),
//   	DigitalSignature: jsii.Boolean(false),
//   	EncipherOnly: jsii.Boolean(false),
//   	KeyAgreement: jsii.Boolean(false),
//   	KeyCertSign: jsii.Boolean(false),
//   	KeyEncipherment: jsii.Boolean(false),
//   	NonRepudiation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html
//
type CfnCertificate_KeyUsageProperty struct {
	// Key can be used to sign CRLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-crlsign
	//
	// Default: - false.
	//
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Key can be used to decipher data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-dataencipherment
	//
	// Default: - false.
	//
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Key can be used only to decipher data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-decipheronly
	//
	// Default: - false.
	//
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Key can be used for digital signing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-digitalsignature
	//
	// Default: - false.
	//
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Key can be used only to encipher data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-encipheronly
	//
	// Default: - false.
	//
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Key can be used in a key-agreement protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-keyagreement
	//
	// Default: - false.
	//
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Key can be used to sign certificates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-keycertsign
	//
	// Default: - false.
	//
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Key can be used to encipher data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-keyencipherment
	//
	// Default: - false.
	//
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Key can be used for non-repudiation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificate-keyusage.html#cfn-acmpca-certificate-keyusage-nonrepudiation
	//
	// Default: - false.
	//
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

