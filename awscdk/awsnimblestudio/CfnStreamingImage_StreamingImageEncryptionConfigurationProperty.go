package awsnimblestudio


// Specifies how a streaming image is encrypted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingImageEncryptionConfigurationProperty := &StreamingImageEncryptionConfigurationProperty{
//   	KeyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	KeyArn: jsii.String("keyArn"),
//   }
//
type CfnStreamingImage_StreamingImageEncryptionConfigurationProperty struct {
	// The type of KMS key that is used to encrypt studio data.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ARN for a KMS key that is used to encrypt studio data.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

