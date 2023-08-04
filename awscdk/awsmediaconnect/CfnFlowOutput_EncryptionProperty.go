package awsmediaconnect


// Information about the encryption of the flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionProperty := &EncryptionProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	Algorithm: jsii.String("algorithm"),
//   	KeyType: jsii.String("keyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encryption.html
//
type CfnFlowOutput_EncryptionProperty struct {
	// The Amazon Resource Name (ARN) of the role that you created during setup (when you set up MediaConnect as a trusted entity).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encryption.html#cfn-mediaconnect-flowoutput-encryption-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the secret that you created in AWS Secrets Manager to store the encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encryption.html#cfn-mediaconnect-flowoutput-encryption-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The type of algorithm that is used for static key encryption (such as aes128, aes192, or aes256).
	//
	// If you are using SPEKE or SRT-password encryption, this property must be left blank.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encryption.html#cfn-mediaconnect-flowoutput-encryption-algorithm
	//
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The type of key that is used for the encryption.
	//
	// If you don't specify a `keyType` value, the service uses the default setting ( `static-key` ). Valid key types are: `static-key` , `speke` , and `srt-password` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-encryption.html#cfn-mediaconnect-flowoutput-encryption-keytype
	//
	// Default: - "static-key".
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

