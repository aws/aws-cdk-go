package awsmediatailor


// Properties for CfnProgramPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnProgramMixinProps := &CfnProgramMixinProps{
//   	AdBreaks: []interface{}{
//   		&AdBreakProperty{
//   			AdBreakMetadata: []interface{}{
//   				&KeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MessageType: jsii.String("messageType"),
//   			OffsetMillis: jsii.Number(123),
//   			Slate: &SlateSourceProperty{
//   				SourceLocationName: jsii.String("sourceLocationName"),
//   				VodSourceName: jsii.String("vodSourceName"),
//   			},
//   			SpliceInsertMessage: &SpliceInsertMessageProperty{
//   				AvailNum: jsii.Number(123),
//   				AvailsExpected: jsii.Number(123),
//   				SpliceEventId: jsii.Number(123),
//   				UniqueProgramId: jsii.Number(123),
//   			},
//   			TimeSignalMessage: &TimeSignalMessageProperty{
//   				SegmentationDescriptors: []interface{}{
//   					&SegmentationDescriptorProperty{
//   						SegmentationEventId: jsii.Number(123),
//   						SegmentationTypeId: jsii.Number(123),
//   						SegmentationUpid: jsii.String("segmentationUpid"),
//   						SegmentationUpidType: jsii.Number(123),
//   						SegmentNum: jsii.Number(123),
//   						SegmentsExpected: jsii.Number(123),
//   						SubSegmentNum: jsii.Number(123),
//   						SubSegmentsExpected: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	AudienceMedia: []interface{}{
//   		&AudienceMediaProperty{
//   			AlternateMedia: []interface{}{
//   				&AlternateMediaProperty{
//   					AdBreaks: []interface{}{
//   						&AdBreakProperty{
//   							AdBreakMetadata: []interface{}{
//   								&KeyValuePairProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MessageType: jsii.String("messageType"),
//   							OffsetMillis: jsii.Number(123),
//   							Slate: &SlateSourceProperty{
//   								SourceLocationName: jsii.String("sourceLocationName"),
//   								VodSourceName: jsii.String("vodSourceName"),
//   							},
//   							SpliceInsertMessage: &SpliceInsertMessageProperty{
//   								AvailNum: jsii.Number(123),
//   								AvailsExpected: jsii.Number(123),
//   								SpliceEventId: jsii.Number(123),
//   								UniqueProgramId: jsii.Number(123),
//   							},
//   							TimeSignalMessage: &TimeSignalMessageProperty{
//   								SegmentationDescriptors: []interface{}{
//   									&SegmentationDescriptorProperty{
//   										SegmentationEventId: jsii.Number(123),
//   										SegmentationTypeId: jsii.Number(123),
//   										SegmentationUpid: jsii.String("segmentationUpid"),
//   										SegmentationUpidType: jsii.Number(123),
//   										SegmentNum: jsii.Number(123),
//   										SegmentsExpected: jsii.Number(123),
//   										SubSegmentNum: jsii.Number(123),
//   										SubSegmentsExpected: jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					ClipRange: &ClipRangeProperty{
//   						EndOffsetMillis: jsii.Number(123),
//   						StartOffsetMillis: jsii.Number(123),
//   					},
//   					DurationMillis: jsii.Number(123),
//   					LiveSourceName: jsii.String("liveSourceName"),
//   					ScheduledStartTimeMillis: jsii.Number(123),
//   					SourceLocationName: jsii.String("sourceLocationName"),
//   					VodSourceName: jsii.String("vodSourceName"),
//   				},
//   			},
//   			Audience: jsii.String("audience"),
//   		},
//   	},
//   	ChannelName: jsii.String("channelName"),
//   	LiveSourceName: jsii.String("liveSourceName"),
//   	ProgramName: jsii.String("programName"),
//   	ScheduleConfiguration: &ScheduleConfigurationProperty{
//   		ClipRange: &ClipRangeProperty{
//   			EndOffsetMillis: jsii.Number(123),
//   			StartOffsetMillis: jsii.Number(123),
//   		},
//   		Transition: &TransitionProperty{
//   			DurationMillis: jsii.Number(123),
//   			RelativePosition: jsii.String("relativePosition"),
//   			RelativeProgram: jsii.String("relativeProgram"),
//   			ScheduledStartTimeMillis: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html
//
type CfnProgramMixinProps struct {
	// The ad break configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-adbreaks
	//
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// The list of AudienceMedia defined in program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-audiencemedia
	//
	AudienceMedia interface{} `field:"optional" json:"audienceMedia" yaml:"audienceMedia"`
	// The name of the channel for this Program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The name of the LiveSource for this Program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-livesourcename
	//
	LiveSourceName *string `field:"optional" json:"liveSourceName" yaml:"liveSourceName"`
	// The name of the Program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-programname
	//
	ProgramName *string `field:"optional" json:"programName" yaml:"programName"`
	// Schedule configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-scheduleconfiguration
	//
	ScheduleConfiguration interface{} `field:"optional" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The name of the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The name that's used to refer to a VOD source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-vodsourcename
	//
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

