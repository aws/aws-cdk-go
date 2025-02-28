package awspcaconnectorad


// Certificate extensions for v2 template schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionsV2Property := &ExtensionsV2Property{
//   	KeyUsage: &KeyUsageProperty{
//   		UsageFlags: &KeyUsageFlagsProperty{
//   			DataEncipherment: jsii.Boolean(false),
//   			DigitalSignature: jsii.Boolean(false),
//   			KeyAgreement: jsii.Boolean(false),
//   			KeyEncipherment: jsii.Boolean(false),
//   			NonRepudiation: jsii.Boolean(false),
//   		},
//
//   		// the properties below are optional
//   		Critical: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	ApplicationPolicies: &ApplicationPoliciesProperty{
//   		Policies: []interface{}{
//   			&ApplicationPolicyProperty{
//   				PolicyObjectIdentifier: jsii.String("policyObjectIdentifier"),
//   				PolicyType: jsii.String("policyType"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Critical: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv2.html
//
type CfnTemplate_ExtensionsV2Property struct {
	// The key usage extension defines the purpose (e.g., encipherment, signature, certificate signing) of the key contained in the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv2.html#cfn-pcaconnectorad-template-extensionsv2-keyusage
	//
	KeyUsage interface{} `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// Application policies specify what the certificate is used for and its purpose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-extensionsv2.html#cfn-pcaconnectorad-template-extensionsv2-applicationpolicies
	//
	ApplicationPolicies interface{} `field:"optional" json:"applicationPolicies" yaml:"applicationPolicies"`
}

