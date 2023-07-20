package awsnimblestudio


// <p>TODO</p>.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html
//
type CfnStreamingImage_StreamingImageEncryptionConfigurationProperty struct {
	// <p/>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html#cfn-nimblestudio-streamingimage-streamingimageencryptionconfiguration-keytype
	//
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// <p>The ARN for a KMS key that is used to encrypt studio data.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html#cfn-nimblestudio-streamingimage-streamingimageencryptionconfiguration-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

