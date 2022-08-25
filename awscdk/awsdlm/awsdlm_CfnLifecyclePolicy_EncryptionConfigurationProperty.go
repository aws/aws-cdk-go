package awsdlm


// Specifies the encryption settings for shared snapshots that are copied across Regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	encrypted: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	cmkArn: jsii.String("cmkArn"),
//   }
//
type CfnLifecyclePolicy_EncryptionConfigurationProperty struct {
	// To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable encryption using this parameter.
	//
	// Copies of encrypted snapshots are encrypted, even if this parameter is false or when encryption by default is not enabled.
	Encrypted interface{} `field:"required" json:"encrypted" yaml:"encrypted"`
	// The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption.
	//
	// If this parameter is not specified, the default KMS key for the account is used.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
}

