package awsivs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRecordingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRecordingConfigurationProps := &CfnRecordingConfigurationProps{
//   	DestinationConfiguration: &DestinationConfigurationProperty{
//   		S3: &S3DestinationConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	RecordingReconnectWindowSeconds: jsii.Number(123),
//   	RenditionConfiguration: &RenditionConfigurationProperty{
//   		Renditions: []*string{
//   			jsii.String("renditions"),
//   		},
//   		RenditionSelection: jsii.String("renditionSelection"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThumbnailConfiguration: &ThumbnailConfigurationProperty{
//   		RecordingMode: jsii.String("recordingMode"),
//   		Resolution: jsii.String("resolution"),
//   		Storage: []*string{
//   			jsii.String("storage"),
//   		},
//   		TargetIntervalSeconds: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html
//
type CfnRecordingConfigurationProps struct {
	// A destination configuration contains information about where recorded video will be stored.
	//
	// See the DestinationConfiguration property type for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html#cfn-ivs-recordingconfiguration-destinationconfiguration
	//
	DestinationConfiguration interface{} `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Recording-configuration name.
	//
	// The value does not need to be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html#cfn-ivs-recordingconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
	//
	// *Default* : `0`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html#cfn-ivs-recordingconfiguration-recordingreconnectwindowseconds
	//
	// Default: - 0.
	//
	RecordingReconnectWindowSeconds *float64 `field:"optional" json:"recordingReconnectWindowSeconds" yaml:"recordingReconnectWindowSeconds"`
	// A rendition configuration describes which renditions should be recorded for a stream.
	//
	// See the RenditionConfiguration property type for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html#cfn-ivs-recordingconfiguration-renditionconfiguration
	//
	RenditionConfiguration interface{} `field:"optional" json:"renditionConfiguration" yaml:"renditionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html#cfn-ivs-recordingconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A thumbnail configuration enables/disables the recording of thumbnails for a live session and controls the interval at which thumbnails are generated for the live session.
	//
	// See the ThumbnailConfiguration property type for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-recordingconfiguration.html#cfn-ivs-recordingconfiguration-thumbnailconfiguration
	//
	ThumbnailConfiguration interface{} `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
}

