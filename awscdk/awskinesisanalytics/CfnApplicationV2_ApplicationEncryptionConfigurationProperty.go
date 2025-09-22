package awskinesisanalytics


// Specifies the configuration to manage encryption at rest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationEncryptionConfigurationProperty := &ApplicationEncryptionConfigurationProperty{
//   	KeyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	KeyId: jsii.String("keyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html
//
type CfnApplicationV2_ApplicationEncryptionConfigurationProperty struct {
	// Specifies the type of key used for encryption at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html#cfn-kinesisanalyticsv2-application-applicationencryptionconfiguration-keytype
	//
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The key ARN, key ID, alias ARN, or alias name of the KMS key used for encryption at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-applicationencryptionconfiguration.html#cfn-kinesisanalyticsv2-application-applicationencryptionconfiguration-keyid
	//
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
}

