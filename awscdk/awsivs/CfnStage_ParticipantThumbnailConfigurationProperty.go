package awsivs


// Object specifying a configuration of thumbnails for recorded video from an individual participant.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   participantThumbnailConfigurationProperty := &ParticipantThumbnailConfigurationProperty{
//   	RecordingMode: jsii.String("recordingMode"),
//   	Storage: []*string{
//   		jsii.String("storage"),
//   	},
//   	TargetIntervalSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-participantthumbnailconfiguration.html
//
type CfnStage_ParticipantThumbnailConfigurationProperty struct {
	// Thumbnail recording mode.
	//
	// Default: `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-participantthumbnailconfiguration.html#cfn-ivs-stage-participantthumbnailconfiguration-recordingmode
	//
	// Default: - "INTERVAL".
	//
	RecordingMode *string `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Indicates the format in which thumbnails are recorded.
	//
	// `SEQUENTIAL` records all generated thumbnails in a serial manner, to the media/thumbnails/high directory. `LATEST` saves the latest thumbnail in media/latest_thumbnail/high/thumb.jpg and overwrites it at the interval specified by `targetIntervalSeconds` . You can enable both `SEQUENTIAL` and `LATEST` . Default: `SEQUENTIAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-participantthumbnailconfiguration.html#cfn-ivs-stage-participantthumbnailconfiguration-storage
	//
	Storage *[]*string `field:"optional" json:"storage" yaml:"storage"`
	// The targeted thumbnail-generation interval in seconds.
	//
	// This is configurable only if `recordingMode` is `INTERVAL` . Default: 60.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-participantthumbnailconfiguration.html#cfn-ivs-stage-participantthumbnailconfiguration-targetintervalseconds
	//
	// Default: - 60.
	//
	TargetIntervalSeconds *float64 `field:"optional" json:"targetIntervalSeconds" yaml:"targetIntervalSeconds"`
}

