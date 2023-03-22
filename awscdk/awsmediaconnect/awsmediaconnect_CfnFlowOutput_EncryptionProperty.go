package awsmediaconnect


// Information about the encryption of the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionProperty := &encryptionProperty{
//   	roleArn: jsii.String("roleArn"),
//   	secretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	algorithm: jsii.String("algorithm"),
//   	keyType: jsii.String("keyType"),
//   }
//
type CfnFlowOutput_EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ).
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

