package awsmediatailor


// A playlist of media to be played instead of the default media on a particular program.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   alternateMediaProperty := &AlternateMediaProperty{
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
//   	ClipRange: &ClipRangeProperty{
//   		EndOffsetMillis: jsii.Number(123),
//   		StartOffsetMillis: jsii.Number(123),
//   	},
//   	DurationMillis: jsii.Number(123),
//   	LiveSourceName: jsii.String("liveSourceName"),
//   	ScheduledStartTimeMillis: jsii.Number(123),
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	VodSourceName: jsii.String("vodSourceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html
//
type CfnProgram_AlternateMediaProperty struct {
	// Ad break configuration parameters defined in AlternateMedia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-adbreaks
	//
	AdBreaks interface{} `field:"optional" json:"adBreaks" yaml:"adBreaks"`
	// Clip range configuration for the VOD source associated with the program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-cliprange
	//
	ClipRange interface{} `field:"optional" json:"clipRange" yaml:"clipRange"`
	// The duration of the alternateMedia in milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-durationmillis
	//
	DurationMillis *float64 `field:"optional" json:"durationMillis" yaml:"durationMillis"`
	// The name of the live source for alternateMedia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-livesourcename
	//
	LiveSourceName *string `field:"optional" json:"liveSourceName" yaml:"liveSourceName"`
	// The date and time that the alternateMedia is scheduled to start, in epoch milliseconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-scheduledstarttimemillis
	//
	ScheduledStartTimeMillis *float64 `field:"optional" json:"scheduledStartTimeMillis" yaml:"scheduledStartTimeMillis"`
	// The name of the source location for alternateMedia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-sourcelocationname
	//
	SourceLocationName *string `field:"optional" json:"sourceLocationName" yaml:"sourceLocationName"`
	// The name of the VOD source for alternateMedia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-alternatemedia.html#cfn-mediatailor-program-alternatemedia-vodsourcename
	//
	VodSourceName *string `field:"optional" json:"vodSourceName" yaml:"vodSourceName"`
}

