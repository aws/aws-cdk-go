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
//   keyUsageProperty := &keyUsageProperty{
//   	crlSign: jsii.Boolean(false),
//   	dataEncipherment: jsii.Boolean(false),
//   	decipherOnly: jsii.Boolean(false),
//   	digitalSignature: jsii.Boolean(false),
//   	encipherOnly: jsii.Boolean(false),
//   	keyAgreement: jsii.Boolean(false),
//   	keyCertSign: jsii.Boolean(false),
//   	keyEncipherment: jsii.Boolean(false),
//   	nonRepudiation: jsii.Boolean(false),
//   }
//
type CfnCertificate_KeyUsageProperty struct {
	// Key can be used to sign CRLs.
	CrlSign interface{} `field:"optional" json:"crlSign" yaml:"crlSign"`
	// Key can be used to decipher data.
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// Key can be used only to decipher data.
	DecipherOnly interface{} `field:"optional" json:"decipherOnly" yaml:"decipherOnly"`
	// Key can be used for digital signing.
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// Key can be used only to encipher data.
	EncipherOnly interface{} `field:"optional" json:"encipherOnly" yaml:"encipherOnly"`
	// Key can be used in a key-agreement protocol.
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Key can be used to sign certificates.
	KeyCertSign interface{} `field:"optional" json:"keyCertSign" yaml:"keyCertSign"`
	// Key can be used to encipher data.
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// Key can be used for non-repudiation.
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

