package awsmediatailor


// An AudienceMedia object contains an Audience and a list of AlternateMedia.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   audienceMediaProperty := &AudienceMediaProperty{
//   	AlternateMedia: []interface{}{
//   		&AlternateMediaProperty{
//   			AdBreaks: []interface{}{
//   				&AdBreakProperty{
//   					AdBreakMetadata: []interface{}{
//   						&KeyValuePairProperty{
//   							Key: jsii.String("key"),
//   							Value: jsii.String("value"),
//   						},
//   					},
//   					MessageType: jsii.String("messageType"),
//   					OffsetMillis: jsii.Number(123),
//   					Slate: &SlateSourceProperty{
//   						SourceLocationName: jsii.String("sourceLocationName"),
//   						VodSourceName: jsii.String("vodSourceName"),
//   					},
//   					SpliceInsertMessage: &SpliceInsertMessageProperty{
//   						AvailNum: jsii.Number(123),
//   						AvailsExpected: jsii.Number(123),
//   						SpliceEventId: jsii.Number(123),
//   						UniqueProgramId: jsii.Number(123),
//   					},
//   					TimeSignalMessage: &TimeSignalMessageProperty{
//   						SegmentationDescriptors: []interface{}{
//   							&SegmentationDescriptorProperty{
//   								SegmentationEventId: jsii.Number(123),
//   								SegmentationTypeId: jsii.Number(123),
//   								SegmentationUpid: jsii.String("segmentationUpid"),
//   								SegmentationUpidType: jsii.Number(123),
//   								SegmentNum: jsii.Number(123),
//   								SegmentsExpected: jsii.Number(123),
//   								SubSegmentNum: jsii.Number(123),
//   								SubSegmentsExpected: jsii.Number(123),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			ClipRange: &ClipRangeProperty{
//   				EndOffsetMillis: jsii.Number(123),
//   				StartOffsetMillis: jsii.Number(123),
//   			},
//   			DurationMillis: jsii.Number(123),
//   			LiveSourceName: jsii.String("liveSourceName"),
//   			ScheduledStartTimeMillis: jsii.Number(123),
//   			SourceLocationName: jsii.String("sourceLocationName"),
//   			VodSourceName: jsii.String("vodSourceName"),
//   		},
//   	},
//   	Audience: jsii.String("audience"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-audiencemedia.html
//
type CfnProgramPropsMixin_AudienceMediaProperty struct {
	// The list of AlternateMedia defined in AudienceMedia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-audiencemedia.html#cfn-mediatailor-program-audiencemedia-alternatemedia
	//
	AlternateMedia interface{} `field:"optional" json:"alternateMedia" yaml:"alternateMedia"`
	// The Audience defined in AudienceMedia.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-audiencemedia.html#cfn-mediatailor-program-audiencemedia-audience
	//
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

