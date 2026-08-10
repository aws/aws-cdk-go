package awshealthlake


// Create the profile by cloning a specific version of an existing profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   existingVersionedProfileSourceProperty := &ExistingVersionedProfileSourceProperty{
//   	ProfileId: jsii.String("profileId"),
//   	Version: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-existingversionedprofilesource.html
//
type CfnDataTransformationProfile_ExistingVersionedProfileSourceProperty struct {
	// The unique identifier of the source profile to clone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-existingversionedprofilesource.html#cfn-healthlake-datatransformationprofile-existingversionedprofilesource-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The version number of the source profile to clone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-datatransformationprofile-existingversionedprofilesource.html#cfn-healthlake-datatransformationprofile-existingversionedprofilesource-version
	//
	Version *float64 `field:"required" json:"version" yaml:"version"`
}

