package mixinsawskinesisanalyticsv2


// Specifies the configuration to manage encryption at rest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationEncryptionConfigurationProperty := &ApplicationEncryptionConfigurationProperty{
//   	KeyId: jsii.String("keyId"),
//   	KeyType: jsii.String("keyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html
//
type CfnApplicationPropsMixin_ApplicationEncryptionConfigurationProperty struct {
	// The key ARN, key ID, alias ARN, or alias name of the KMS key used for encryption at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html#cfn-kinesisanalyticsv2-application-applicationencryptionconfiguration-keyid
	//
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Specifies the type of key used for encryption at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html#cfn-kinesisanalyticsv2-application-applicationencryptionconfiguration-keytype
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

