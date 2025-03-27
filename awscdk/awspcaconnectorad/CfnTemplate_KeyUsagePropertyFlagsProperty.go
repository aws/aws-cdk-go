package awspcaconnectorad


// Specifies key usage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyUsagePropertyFlagsProperty := &KeyUsagePropertyFlagsProperty{
//   	Decrypt: jsii.Boolean(false),
//   	KeyAgreement: jsii.Boolean(false),
//   	Sign: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html
//
type CfnTemplate_KeyUsagePropertyFlagsProperty struct {
	// Allows key for encryption and decryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html#cfn-pcaconnectorad-template-keyusagepropertyflags-decrypt
	//
	Decrypt interface{} `field:"optional" json:"decrypt" yaml:"decrypt"`
	// Allows key exchange without encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html#cfn-pcaconnectorad-template-keyusagepropertyflags-keyagreement
	//
	KeyAgreement interface{} `field:"optional" json:"keyAgreement" yaml:"keyAgreement"`
	// Allow key use for digital signature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusagepropertyflags.html#cfn-pcaconnectorad-template-keyusagepropertyflags-sign
	//
	Sign interface{} `field:"optional" json:"sign" yaml:"sign"`
}

