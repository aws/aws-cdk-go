package awsnimblestudio


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingImageEncryptionConfigurationProperty := &streamingImageEncryptionConfigurationProperty{
//   	keyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	keyArn: jsii.String("keyArn"),
//   }
//
type CfnStreamingImage_StreamingImageEncryptionConfigurationProperty struct {
	// `CfnStreamingImage.StreamingImageEncryptionConfigurationProperty.KeyType`.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// `CfnStreamingImage.StreamingImageEncryptionConfigurationProperty.KeyArn`.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

