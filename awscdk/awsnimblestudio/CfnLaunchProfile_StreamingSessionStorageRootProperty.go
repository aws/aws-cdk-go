package awsnimblestudio


// The upload storage root location (folder) on streaming workstations where files are uploaded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingSessionStorageRootProperty := &StreamingSessionStorageRootProperty{
//   	Linux: jsii.String("linux"),
//   	Windows: jsii.String("windows"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamingsessionstorageroot.html
//
type CfnLaunchProfile_StreamingSessionStorageRootProperty struct {
	// The folder path in Linux workstations where files are uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamingsessionstorageroot.html#cfn-nimblestudio-launchprofile-streamingsessionstorageroot-linux
	//
	Linux *string `field:"optional" json:"linux" yaml:"linux"`
	// The folder path in Windows workstations where files are uploaded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamingsessionstorageroot.html#cfn-nimblestudio-launchprofile-streamingsessionstorageroot-windows
	//
	Windows *string `field:"optional" json:"windows" yaml:"windows"`
}

