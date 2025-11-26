package previewawsnimblestudiomixins


// <p>Configuration of the encryption method that is used for the studio.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   studioEncryptionConfigurationProperty := &StudioEncryptionConfigurationProperty{
//   	KeyArn: jsii.String("keyArn"),
//   	KeyType: jsii.String("keyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studio-studioencryptionconfiguration.html
//
type CfnStudioPropsMixin_StudioEncryptionConfigurationProperty struct {
	// <p>The ARN for a KMS key that is used to encrypt studio data.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studio-studioencryptionconfiguration.html#cfn-nimblestudio-studio-studioencryptionconfiguration-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// <p>The type of KMS key that is used to encrypt studio data.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studio-studioencryptionconfiguration.html#cfn-nimblestudio-studio-studioencryptionconfiguration-keytype
	//
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

