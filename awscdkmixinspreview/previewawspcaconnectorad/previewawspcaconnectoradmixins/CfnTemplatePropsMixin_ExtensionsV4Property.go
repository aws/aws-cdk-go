package previewawspcaconnectoradmixins


// Certificate extensions for v4 template schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   extensionsV4Property := &ExtensionsV4Property{
//   	ApplicationPolicies: &ApplicationPoliciesProperty{
//   		Critical: jsii.Boolean(false),
//   		Policies: []interface{}{
//   			&ApplicationPolicyProperty{
//   				PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   				PolicyType: jsii.String("policyType"),
//   			},
//   		},
//   	},
//   	KeyUsage: &KeyUsageProperty{
//   		Critical: jsii.Boolean(false),
//   		UsageFlags: &KeyUsageFlagsProperty{
//   			DataEncipherment: jsii.Boolean(false),
//   			DigitalSignature: jsii.Boolean(false),
//   			KeyAgreement: jsii.Boolean(false),
//   			KeyEncipherment: jsii.Boolean(false),
//   			NonRepudiation: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv4.html
//
type CfnTemplatePropsMixin_ExtensionsV4Property struct {
	// Application policies specify what the certificate is used for and its purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv4.html#cfn-pcaconnectorad-template-extensionsv4-applicationpolicies
	//
	ApplicationPolicies interface{} `field:"optional" json:"applicationPolicies" yaml:"applicationPolicies"`
	// The key usage extension defines the purpose (e.g., encipherment, signature) of the key contained in the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv4.html#cfn-pcaconnectorad-template-extensionsv4-keyusage
	//
	KeyUsage interface{} `field:"optional" json:"keyUsage" yaml:"keyUsage"`
}

