package awsnimblestudio


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamingsessionstorageroot.html#cfn-nimblestudio-launchprofile-streamingsessionstorageroot-linux
	//
	Linux *string `field:"optional" json:"linux" yaml:"linux"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-nimblestudio-launchprofile-streamingsessionstorageroot.html#cfn-nimblestudio-launchprofile-streamingsessionstorageroot-windows
	//
	Windows *string `field:"optional" json:"windows" yaml:"windows"`
}

