package awsmediatailor


// The SCTE-35 time_signal message configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeSignalMessageProperty := &TimeSignalMessageProperty{
//   	SegmentationDescriptors: []interface{}{
//   		&SegmentationDescriptorProperty{
//   			SegmentationEventId: jsii.Number(123),
//   			SegmentationTypeId: jsii.Number(123),
//   			SegmentationUpid: jsii.String("segmentationUpid"),
//   			SegmentationUpidType: jsii.Number(123),
//   			SegmentNum: jsii.Number(123),
//   			SegmentsExpected: jsii.Number(123),
//   			SubSegmentNum: jsii.Number(123),
//   			SubSegmentsExpected: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-timesignalmessage.html
//
type CfnProgram_TimeSignalMessageProperty struct {
	// The configurations for the SCTE-35 segmentation_descriptor message(s).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-timesignalmessage.html#cfn-mediatailor-program-timesignalmessage-segmentationdescriptors
	//
	SegmentationDescriptors interface{} `field:"optional" json:"segmentationDescriptors" yaml:"segmentationDescriptors"`
}

