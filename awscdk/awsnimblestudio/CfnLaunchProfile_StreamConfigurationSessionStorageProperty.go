package awsnimblestudio


// The configuration for a streaming session’s upload storage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamConfigurationSessionStorageProperty := &StreamConfigurationSessionStorageProperty{
//   	Mode: []*string{
//   		jsii.String("mode"),
//   	},
//
//   	// the properties below are optional
//   	Root: &StreamingSessionStorageRootProperty{
//   		Linux: jsii.String("linux"),
//   		Windows: jsii.String("windows"),
//   	},
//   }
//
type CfnLaunchProfile_StreamConfigurationSessionStorageProperty struct {
	// Allows artists to upload files to their workstations.
	//
	// The only valid option is `UPLOAD` .
	Mode *[]*string `field:"required" json:"mode" yaml:"mode"`
	// The configuration for the upload storage root of the streaming session.
	Root interface{} `field:"optional" json:"root" yaml:"root"`
}

