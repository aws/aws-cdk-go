package previewawsnimblestudiomixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamingImageEncryptionConfigurationProperty := &StreamingImageEncryptionConfigurationProperty{
//   	KeyArn: jsii.String("keyArn"),
//   	KeyType: jsii.String("keyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html
//
type CfnStreamingImagePropsMixin_StreamingImageEncryptionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html#cfn-nimblestudio-streamingimage-streamingimageencryptionconfiguration-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-streamingimage-streamingimageencryptionconfiguration.html#cfn-nimblestudio-streamingimage-streamingimageencryptionconfiguration-keytype
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

