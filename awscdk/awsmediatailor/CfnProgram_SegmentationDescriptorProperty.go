package awsmediatailor


// The segmentation_descriptor message configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentationDescriptorProperty := &SegmentationDescriptorProperty{
//   	SegmentationEventId: jsii.Number(123),
//   	SegmentationTypeId: jsii.Number(123),
//   	SegmentationUpid: jsii.String("segmentationUpid"),
//   	SegmentationUpidType: jsii.Number(123),
//   	SegmentNum: jsii.Number(123),
//   	SegmentsExpected: jsii.Number(123),
//   	SubSegmentNum: jsii.Number(123),
//   	SubSegmentsExpected: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html
//
type CfnProgram_SegmentationDescriptorProperty struct {
	// The Event Identifier to assign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-segmentationeventid
	//
	SegmentationEventId *float64 `field:"optional" json:"segmentationEventId" yaml:"segmentationEventId"`
	// The Type Identifier to assign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-segmentationtypeid
	//
	SegmentationTypeId *float64 `field:"optional" json:"segmentationTypeId" yaml:"segmentationTypeId"`
	// The Upid to assign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-segmentationupid
	//
	SegmentationUpid *string `field:"optional" json:"segmentationUpid" yaml:"segmentationUpid"`
	// The Upid Type to assign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-segmentationupidtype
	//
	SegmentationUpidType *float64 `field:"optional" json:"segmentationUpidType" yaml:"segmentationUpidType"`
	// The segment number to assign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-segmentnum
	//
	SegmentNum *float64 `field:"optional" json:"segmentNum" yaml:"segmentNum"`
	// The number of segments expected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-segmentsexpected
	//
	SegmentsExpected *float64 `field:"optional" json:"segmentsExpected" yaml:"segmentsExpected"`
	// The sub-segment number to assign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-subsegmentnum
	//
	SubSegmentNum *float64 `field:"optional" json:"subSegmentNum" yaml:"subSegmentNum"`
	// The number of sub-segments expected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-segmentationdescriptor.html#cfn-mediatailor-program-segmentationdescriptor-subsegmentsexpected
	//
	SubSegmentsExpected *float64 `field:"optional" json:"subSegmentsExpected" yaml:"subSegmentsExpected"`
}

