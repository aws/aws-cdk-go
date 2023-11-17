package awssagemaker


// Default storage and volume settings for a space.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacestoragesettings.html
//
type CfnDomain_DefaultSpaceStorageSettingsProperty struct {
	// Properties related to the Amazon Elastic Block Store volume.
	//
	// Must be provided if storage type is Amazon EBS and must not be provided if storage type is not Amazon EBS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacestoragesettings.html#cfn-sagemaker-domain-defaultspacestoragesettings-defaultebsstoragesettings
	//
	DefaultEbsStorageSettings interface{} `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}

