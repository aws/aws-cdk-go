package previewawssagemakermixins


// A collection of default EBS storage settings that apply to spaces created within a domain or user profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   defaultEbsStorageSettingsProperty := &DefaultEbsStorageSettingsProperty{
//   	DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   	MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-defaultebsstoragesettings.html
//
type CfnUserProfilePropsMixin_DefaultEbsStorageSettingsProperty struct {
	// The default size of the EBS storage volume for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-defaultebsstoragesettings.html#cfn-sagemaker-userprofile-defaultebsstoragesettings-defaultebsvolumesizeingb
	//
	DefaultEbsVolumeSizeInGb *float64 `field:"optional" json:"defaultEbsVolumeSizeInGb" yaml:"defaultEbsVolumeSizeInGb"`
	// The maximum size of the EBS storage volume for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-defaultebsstoragesettings.html#cfn-sagemaker-userprofile-defaultebsstoragesettings-maximumebsvolumesizeingb
	//
	MaximumEbsVolumeSizeInGb *float64 `field:"optional" json:"maximumEbsVolumeSizeInGb" yaml:"maximumEbsVolumeSizeInGb"`
}

