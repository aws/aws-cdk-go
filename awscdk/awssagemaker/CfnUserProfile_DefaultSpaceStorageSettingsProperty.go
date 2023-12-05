package awssagemaker


// Default storage settings for a space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultSpaceStorageSettingsProperty := &DefaultSpaceStorageSettingsProperty{
//   	DefaultEbsStorageSettings: &DefaultEbsStorageSettingsProperty{
//   		DefaultEbsVolumeSizeInGb: jsii.Number(123),
//   		MaximumEbsVolumeSizeInGb: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-defaultspacestoragesettings.html
//
type CfnUserProfile_DefaultSpaceStorageSettingsProperty struct {
	// Properties related to the Amazon Elastic Block Store volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-defaultspacestoragesettings.html#cfn-sagemaker-userprofile-defaultspacestoragesettings-defaultebsstoragesettings
	//
	DefaultEbsStorageSettings interface{} `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}

