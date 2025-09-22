package awspcaconnectorad


// The key usage property defines the purpose of the private key contained in the certificate.
//
// You can specify specific purposes using property flags or all by using property type ALL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyUsagePropertyProperty := &KeyUsagePropertyProperty{
//   	PropertyFlags: &KeyUsagePropertyFlagsProperty{
//   		Decrypt: jsii.Boolean(false),
//   		KeyAgreement: jsii.Boolean(false),
//   		Sign: jsii.Boolean(false),
//   	},
//   	PropertyType: jsii.String("propertyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageproperty.html
//
type CfnTemplate_KeyUsagePropertyProperty struct {
	// You can specify key usage for encryption, key agreement, and signature.
	//
	// You can use property flags or property type but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageproperty.html#cfn-pcaconnectorad-template-keyusageproperty-propertyflags
	//
	PropertyFlags interface{} `field:"optional" json:"propertyFlags" yaml:"propertyFlags"`
	// You can specify all key usages using property type ALL.
	//
	// You can use property type or property flags but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcaconnectorad-template-keyusageproperty.html#cfn-pcaconnectorad-template-keyusageproperty-propertytype
	//
	PropertyType *string `field:"optional" json:"propertyType" yaml:"propertyType"`
}

