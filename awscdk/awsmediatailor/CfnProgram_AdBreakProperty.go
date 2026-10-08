package awsmediatailor


// Ad break configuration parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adBreakProperty := &AdBreakProperty{
//   	OffsetMillis: jsii.Number(123),
//
//   	// the properties below are optional
//   	AdBreakMetadata: []interface{}{
//   		&KeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	MessageType: jsii.String("messageType"),
//   	Slate: &SlateSourceProperty{
//   		SourceLocationName: jsii.String("sourceLocationName"),
//   		VodSourceName: jsii.String("vodSourceName"),
//   	},
//   	SpliceInsertMessage: &SpliceInsertMessageProperty{
//   		AvailNum: jsii.Number(123),
//   		AvailsExpected: jsii.Number(123),
//   		SpliceEventId: jsii.Number(123),
//   		UniqueProgramId: jsii.Number(123),
//   	},
//   	TimeSignalMessage: &TimeSignalMessageProperty{
//   		SegmentationDescriptors: []interface{}{
//   			&SegmentationDescriptorProperty{
//   				SegmentationEventId: jsii.Number(123),
//   				SegmentationTypeId: jsii.Number(123),
//   				SegmentationUpid: jsii.String("segmentationUpid"),
//   				SegmentationUpidType: jsii.Number(123),
//   				SegmentNum: jsii.Number(123),
//   				SegmentsExpected: jsii.Number(123),
//   				SubSegmentNum: jsii.Number(123),
//   				SubSegmentsExpected: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html
//
type CfnProgram_AdBreakProperty struct {
	// How long (in milliseconds) after the beginning of the program that an ad starts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html#cfn-mediatailor-program-adbreak-offsetmillis
	//
	OffsetMillis *float64 `field:"required" json:"offsetMillis" yaml:"offsetMillis"`
	// Defines a list of key/value pairs that MediaTailor generates within the EXT-X-ASSET tag for SCTE35_ENHANCED output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html#cfn-mediatailor-program-adbreak-adbreakmetadata
	//
	AdBreakMetadata interface{} `field:"optional" json:"adBreakMetadata" yaml:"adBreakMetadata"`
	// The SCTE-35 ad insertion type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html#cfn-mediatailor-program-adbreak-messagetype
	//
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
	// Slate VOD source configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html#cfn-mediatailor-program-adbreak-slate
	//
	Slate interface{} `field:"optional" json:"slate" yaml:"slate"`
	// Splice insert message configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html#cfn-mediatailor-program-adbreak-spliceinsertmessage
	//
	SpliceInsertMessage interface{} `field:"optional" json:"spliceInsertMessage" yaml:"spliceInsertMessage"`
	// The SCTE-35 time_signal message configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-adbreak.html#cfn-mediatailor-program-adbreak-timesignalmessage
	//
	TimeSignalMessage interface{} `field:"optional" json:"timeSignalMessage" yaml:"timeSignalMessage"`
}

