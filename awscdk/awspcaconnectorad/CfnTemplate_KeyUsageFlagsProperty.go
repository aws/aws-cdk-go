package awspcaconnectorad


// The key usage flags represent the purpose (e.g., encipherment, signature) of the key contained in the certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyUsageFlagsProperty := &KeyUsageFlagsProperty{
//   	DataEncipherment: jsii.Boolean(false),
//   	DigitalSignature: jsii.Boolean(false),
//   	KeyAgreement: jsii.Boolean(false),
//   	KeyEncipherment: jsii.Boolean(false),
//   	NonRepudiation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageflags.html
//
type CfnTemplate_KeyUsageFlagsProperty struct {
	// DataEncipherment is asserted when the subject public key is used for directly enciphering raw user data without the use of an intermediate symmetric cipher.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageflags.html#cfn-pcaconnectorad-template-keyusageflags-dataencipherment
	//
	DataEncipherment interface{} `field:"optional" json:"dataEncipherment" yaml:"dataEncipherment"`
	// The digitalSignature is asserted when the subject public key is used for verifying digital signatures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageflags.html#cfn-pcaconnectorad-template-keyusageflags-digitalsignature
	//
	DigitalSignature interface{} `field:"optional" json:"digitalSignature" yaml:"digitalSignature"`
	// KeyAgreement is asserted when the subject public key is used for key agreement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageflags.html#cfn-pcaconnectorad-template-keyusageflags-keyagreement
	//
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// KeyEncipherment is asserted when the subject public key is used for enciphering private or secret keys, i.e., for key transport.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageflags.html#cfn-pcaconnectorad-template-keyusageflags-keyencipherment
	//
	KeyEncipherment interface{} `field:"optional" json:"keyEncipherment" yaml:"keyEncipherment"`
	// NonRepudiation is asserted when the subject public key is used to verify digital signatures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageflags.html#cfn-pcaconnectorad-template-keyusageflags-nonrepudiation
	//
	NonRepudiation interface{} `field:"optional" json:"nonRepudiation" yaml:"nonRepudiation"`
}

