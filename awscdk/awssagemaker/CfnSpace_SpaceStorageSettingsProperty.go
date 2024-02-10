package awssagemaker


// The storage settings for a private space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceStorageSettingsProperty := &SpaceStorageSettingsProperty{
//   	EbsStorageSettings: &EbsStorageSettingsProperty{
//   		EbsVolumeSizeInGb: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacestoragesettings.html
//
type CfnSpace_SpaceStorageSettingsProperty struct {
	// A collection of EBS storage settings for a private space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacestoragesettings.html#cfn-sagemaker-space-spacestoragesettings-ebsstoragesettings
	//
	EbsStorageSettings interface{} `field:"optional" json:"ebsStorageSettings" yaml:"ebsStorageSettings"`
}

