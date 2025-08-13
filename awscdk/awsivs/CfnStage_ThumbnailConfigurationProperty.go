package awsivs


// A complex type that allows you to enable/disable the recording of thumbnails for individual participant recording and modify the interval at which thumbnails are generated for the live session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thumbnailConfigurationProperty := &ThumbnailConfigurationProperty{
//   	ParticipantThumbnailConfiguration: &ParticipantThumbnailConfigurationProperty{
//   		RecordingMode: jsii.String("recordingMode"),
//   		Storage: []*string{
//   			jsii.String("storage"),
//   		},
//   		TargetIntervalSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-thumbnailconfiguration.html
//
type CfnStage_ThumbnailConfigurationProperty struct {
	// An object representing a configuration of thumbnails for recorded video from an individual participant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-thumbnailconfiguration.html#cfn-ivs-stage-thumbnailconfiguration-participantthumbnailconfiguration
	//
	ParticipantThumbnailConfiguration interface{} `field:"optional" json:"participantThumbnailConfiguration" yaml:"participantThumbnailConfiguration"`
}

