package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   transcriptFilterProperty := &TranscriptFilterProperty{
//   	AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   		EndTime: jsii.Number(123),
//   		First: jsii.Number(123),
//   		Last: jsii.Number(123),
//   		StartTime: jsii.Number(123),
//   	},
//   	Negate: jsii.Boolean(false),
//   	ParticipantRole: jsii.String("participantRole"),
//   	RelativeTimeRange: &RelativeTimeRangeProperty{
//   		EndPercentage: jsii.Number(123),
//   		First: jsii.Number(123),
//   		Last: jsii.Number(123),
//   		StartPercentage: jsii.Number(123),
//   	},
//   	Targets: []*string{
//   		jsii.String("targets"),
//   	},
//   	TranscriptFilterType: jsii.String("transcriptFilterType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html
//
type CfnCallAnalyticsCategoryPropsMixin_TranscriptFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html#cfn-transcribe-callanalyticscategory-transcriptfilter-absolutetimerange
	//
	AbsoluteTimeRange interface{} `field:"optional" json:"absoluteTimeRange" yaml:"absoluteTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html#cfn-transcribe-callanalyticscategory-transcriptfilter-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html#cfn-transcribe-callanalyticscategory-transcriptfilter-participantrole
	//
	ParticipantRole *string `field:"optional" json:"participantRole" yaml:"participantRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html#cfn-transcribe-callanalyticscategory-transcriptfilter-relativetimerange
	//
	RelativeTimeRange interface{} `field:"optional" json:"relativeTimeRange" yaml:"relativeTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html#cfn-transcribe-callanalyticscategory-transcriptfilter-targets
	//
	Targets *[]*string `field:"optional" json:"targets" yaml:"targets"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-transcriptfilter.html#cfn-transcribe-callanalyticscategory-transcriptfilter-transcriptfiltertype
	//
	TranscriptFilterType *string `field:"optional" json:"transcriptFilterType" yaml:"transcriptFilterType"`
}

