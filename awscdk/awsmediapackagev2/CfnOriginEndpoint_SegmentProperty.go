package awsmediapackagev2


// The segment configuration, including the segment name, duration, and other configuration values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentProperty := &SegmentProperty{
//   	Encryption: &EncryptionProperty{
//   		EncryptionMethod: &EncryptionMethodProperty{
//   			CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   			IsmEncryptionMethod: jsii.String("ismEncryptionMethod"),
//   			TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   		},
//   		SpekeKeyProvider: &SpekeKeyProviderProperty{
//   			DrmSystems: []*string{
//   				jsii.String("drmSystems"),
//   			},
//   			EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   			ResourceId: jsii.String("resourceId"),
//   			RoleArn: jsii.String("roleArn"),
//   			Url: jsii.String("url"),
//   		},
//
//   		// the properties below are optional
//   		CmafExcludeSegmentDrmMetadata: jsii.Boolean(false),
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		KeyRotationIntervalSeconds: jsii.Number(123),
//   	},
//   	IncludeIframeOnlyStreams: jsii.Boolean(false),
//   	Scte: &ScteProperty{
//   		ScteFilter: []*string{
//   			jsii.String("scteFilter"),
//   		},
//   		ScteInSegments: jsii.String("scteInSegments"),
//   	},
//   	SegmentDurationSeconds: jsii.Number(123),
//   	SegmentName: jsii.String("segmentName"),
//   	TsIncludeDvbSubtitles: jsii.Boolean(false),
//   	TsUseAudioRenditionGroup: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html
//
type CfnOriginEndpoint_SegmentProperty struct {
	// Whether to use encryption for the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// Whether the segment includes I-frame-only streams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-includeiframeonlystreams
	//
	IncludeIframeOnlyStreams interface{} `field:"optional" json:"includeIframeOnlyStreams" yaml:"includeIframeOnlyStreams"`
	// The SCTE-35 configuration associated with the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-scte
	//
	Scte interface{} `field:"optional" json:"scte" yaml:"scte"`
	// The duration of the segment, in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-segmentdurationseconds
	//
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// The name of the segment associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-segmentname
	//
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// Whether the segment includes DVB subtitles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-tsincludedvbsubtitles
	//
	TsIncludeDvbSubtitles interface{} `field:"optional" json:"tsIncludeDvbSubtitles" yaml:"tsIncludeDvbSubtitles"`
	// Whether the segment is an audio rendition group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-segment.html#cfn-mediapackagev2-originendpoint-segment-tsuseaudiorenditiongroup
	//
	TsUseAudioRenditionGroup interface{} `field:"optional" json:"tsUseAudioRenditionGroup" yaml:"tsUseAudioRenditionGroup"`
}

