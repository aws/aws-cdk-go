package awsmediatailor


// Properties for defining a `CfnProgram`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProgramProps := &CfnProgramProps{
//   	ChannelName: jsii.String("channelName"),
//   	ProgramName: jsii.String("programName"),
//   	SourceLocationName: jsii.String("sourceLocationName"),
//
//   	// the properties below are optional
//   	AdBreaks: []interface{}{
//   		&AdBreakProperty{
//   			OffsetMillis: jsii.Number(123),
//
//   			// the properties below are optional
//   			AdBreakMetadata: []interface{}{
//   				&KeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MessageType: jsii.String("messageType"),
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
//   							OffsetMillis: jsii.Number(123),
//
//   							// the properties below are optional
//   							AdBreakMetadata: []interface{}{
//   								&KeyValuePairProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MessageType: jsii.String("messageType"),
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
//   	LiveSourceName: jsii.String("liveSourceName"),
//   	ScheduleConfiguration: &ScheduleConfigurationProperty{
//   		Transition: &TransitionProperty{
//   			RelativePosition: jsii.String("relativePosition"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			DurationMillis: jsii.Number(123),
//   			RelativeProgram: jsii.String("relativeProgram"),
//   			ScheduledStartTimeMillis: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		ClipRange: &ClipRangeProperty{
//   			EndOffsetMillis: jsii.Number(123),
//   			StartOffsetMillis: jsii.Number(123),
//   		},
//   	},
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html
//
type CfnProgramProps struct {
	// The name of the channel for this Program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The name of the Program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-programname
	//
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
	// The name of the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-sourcelocationname
	//
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The ad break configuration settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-adbreaks
	//
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// The list of AudienceMedia defined in program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-audiencemedia
	//
	AudienceMedia interface{} `field:"optional" json:"audienceMedia" yaml:"audienceMedia"`
	// The name of the LiveSource for this Program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-livesourcename
	//
	LiveSourceName *string `field:"optional" json:"liveSourceName" yaml:"liveSourceName"`
	// Schedule configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-scheduleconfiguration
	//
	ScheduleConfiguration interface{} `field:"optional" json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The name that's used to refer to a VOD source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html#cfn-mediatailor-program-vodsourcename
	//
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

