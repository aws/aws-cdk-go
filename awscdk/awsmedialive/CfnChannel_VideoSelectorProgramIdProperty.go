package awsmedialive


// Used to extract video by the program ID.
//
// The parent of this entity is VideoSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoSelectorProgramIdProperty := &VideoSelectorProgramIdProperty{
//   	ProgramId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorprogramid.html
//
type CfnChannel_VideoSelectorProgramIdProperty struct {
	// Selects a specific program from within a multi-program transport stream.
	//
	// If the program doesn't exist, MediaLive selects the first program within the transport stream by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorprogramid.html#cfn-medialive-channel-videoselectorprogramid-programid
	//
	ProgramId *float64 `field:"optional" json:"programId" yaml:"programId"`
}

