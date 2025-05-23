package awsdlm


// *[Event-based policies only]* Specifies the encryption settings for cross-Region snapshot copies created by event-based policies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	Encrypted: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	CmkArn: jsii.String("cmkArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-encryptionconfiguration.html
//
type CfnLifecyclePolicy_EncryptionConfigurationProperty struct {
	// To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or when encryption by default is not enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-encryptionconfiguration.html#cfn-dlm-lifecyclepolicy-encryptionconfiguration-encrypted
	//
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-encryptionconfiguration.html#cfn-dlm-lifecyclepolicy-encryptionconfiguration-cmkarn
	//
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
}

