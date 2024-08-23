package awsivs


// Configuration object for individual participant recording, to attach to the new stage.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoParticipantRecordingConfigurationProperty := &AutoParticipantRecordingConfigurationProperty{
//   	StorageConfigurationArn: jsii.String("storageConfigurationArn"),
//
//   	// the properties below are optional
//   	MediaTypes: []*string{
//   		jsii.String("mediaTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html
//
type CfnStage_AutoParticipantRecordingConfigurationProperty struct {
	// ARN of the StorageConfiguration resource to use for individual participant recording.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-storageconfigurationarn
	//
	StorageConfigurationArn *string `field:"required" json:"storageConfigurationArn" yaml:"storageConfigurationArn"`
	// Types of media to be recorded.
	//
	// Default: AUDIO_VIDEO.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-autoparticipantrecordingconfiguration.html#cfn-ivs-stage-autoparticipantrecordingconfiguration-mediatypes
	//
	MediaTypes *[]*string `field:"optional" json:"mediaTypes" yaml:"mediaTypes"`
}

