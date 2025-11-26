package previewawspcaconnectoradmixins


// The key usage extension defines the purpose (e.g., encipherment, signature) of the key contained in the certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   keyUsageProperty := &KeyUsageProperty{
//   	Critical: jsii.Boolean(false),
//   	UsageFlags: &KeyUsageFlagsProperty{
//   		DataEncipherment: jsii.Boolean(false),
//   		DigitalSignature: jsii.Boolean(false),
//   		KeyAgreement: jsii.Boolean(false),
//   		KeyEncipherment: jsii.Boolean(false),
//   		NonRepudiation: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusage.html
//
type CfnTemplatePropsMixin_KeyUsageProperty struct {
	// Sets the key usage extension to critical.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusage.html#cfn-pcaconnectorad-template-keyusage-critical
	//
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
	// The key usage flags represent the purpose (e.g., encipherment, signature) of the key contained in the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusage.html#cfn-pcaconnectorad-template-keyusage-usageflags
	//
	UsageFlags interface{} `field:"optional" json:"usageFlags" yaml:"usageFlags"`
}

