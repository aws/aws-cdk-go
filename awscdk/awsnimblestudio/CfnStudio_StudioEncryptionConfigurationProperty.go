package awsnimblestudio


// Configuration of the encryption method that is used for the studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioEncryptionConfigurationProperty := &StudioEncryptionConfigurationProperty{
//   	KeyType: jsii.String("keyType"),
//
//   	// the properties below are optional
//   	KeyArn: jsii.String("keyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studio-studioencryptionconfiguration.html
//
type CfnStudio_StudioEncryptionConfigurationProperty struct {
	// The type of KMS key that is used to encrypt studio data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studio-studioencryptionconfiguration.html#cfn-nimblestudio-studio-studioencryptionconfiguration-keytype
	//
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// The ARN for a KMS key that is used to encrypt studio data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-studio-studioencryptionconfiguration.html#cfn-nimblestudio-studio-studioencryptionconfiguration-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

