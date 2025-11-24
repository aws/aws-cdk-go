package mixinsawsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStagePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStageMixinProps := &CfnStageMixinProps{
//   	AutoParticipantRecordingConfiguration: &AutoParticipantRecordingConfigurationProperty{
//   		HlsConfiguration: &HlsConfigurationProperty{
//   			ParticipantRecordingHlsConfiguration: &ParticipantRecordingHlsConfigurationProperty{
//   				TargetSegmentDurationSeconds: jsii.Number(123),
//   			},
//   		},
//   		MediaTypes: []*string{
//   			jsii.String("mediaTypes"),
//   		},
//   		RecordingReconnectWindowSeconds: jsii.Number(123),
//   		StorageConfigurationArn: jsii.String("storageConfigurationArn"),
//   		ThumbnailConfiguration: &ThumbnailConfigurationProperty{
//   			ParticipantThumbnailConfiguration: &ParticipantThumbnailConfigurationProperty{
//   				RecordingMode: jsii.String("recordingMode"),
//   				Storage: []*string{
//   					jsii.String("storage"),
//   				},
//   				TargetIntervalSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-stage.html
//
type CfnStageMixinProps struct {
	// Configuration object for individual participant recording.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-stage.html#cfn-ivs-stage-autoparticipantrecordingconfiguration
	//
	AutoParticipantRecordingConfiguration interface{} `field:"optional" json:"autoParticipantRecordingConfiguration" yaml:"autoParticipantRecordingConfiguration"`
	// Stage name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-stage.html#cfn-ivs-stage-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ivs-stage-tag.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-stage.html#cfn-ivs-stage-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

