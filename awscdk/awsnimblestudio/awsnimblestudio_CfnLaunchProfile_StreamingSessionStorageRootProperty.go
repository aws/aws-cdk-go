package awsnimblestudio


// The upload storage root location (folder) on streaming workstations where files are uploaded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingSessionStorageRootProperty := &streamingSessionStorageRootProperty{
//   	linux: jsii.String("linux"),
//   	windows: jsii.String("windows"),
//   }
//
type CfnLaunchProfile_StreamingSessionStorageRootProperty struct {
	// The folder path in Linux workstations where files are uploaded.
	Linux *string `field:"optional" json:"linux" yaml:"linux"`
	// The folder path in Windows workstations where files are uploaded.
	Windows *string `field:"optional" json:"windows" yaml:"windows"`
}

