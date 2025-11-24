package mixinsawssagemaker


// A collection of EBS storage settings that apply to both private and shared spaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ebsStorageSettingsProperty := &EbsStorageSettingsProperty{
//   	EbsVolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-ebsstoragesettings.html
//
type CfnSpacePropsMixin_EbsStorageSettingsProperty struct {
	// The size of an EBS storage volume for a space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-ebsstoragesettings.html#cfn-sagemaker-space-ebsstoragesettings-ebsvolumesizeingb
	//
	EbsVolumeSizeInGb *float64 `field:"optional" json:"ebsVolumeSizeInGb" yaml:"ebsVolumeSizeInGb"`
}

