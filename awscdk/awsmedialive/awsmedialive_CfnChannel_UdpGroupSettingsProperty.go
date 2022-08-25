package awsmedialive


// The configuration of a UDP output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   udpGroupSettingsProperty := &udpGroupSettingsProperty{
//   	inputLossAction: jsii.String("inputLossAction"),
//   	timedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	timedMetadataId3Period: jsii.Number(123),
//   }
//
type CfnChannel_UdpGroupSettingsProperty struct {
	// Specifies the behavior of the last resort when the input video is lost, and no more backup inputs are available.
	//
	// When dropTs is selected, the entire transport stream stops emitting. When dropProgram is selected, the program can be dropped from the transport stream (and replaced with null packets to meet the TS bitrate requirement). Or when emitProgram is selected, the transport stream continues to be produced normally with repeat frames, black frames, or slate frames substituted for the absent input video.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Indicates the ID3 frame that has the timecode.
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval in seconds.
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
}

