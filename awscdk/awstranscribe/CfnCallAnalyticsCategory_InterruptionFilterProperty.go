package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   interruptionFilterProperty := &InterruptionFilterProperty{
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
//   	Threshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-interruptionfilter.html
//
type CfnCallAnalyticsCategory_InterruptionFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-interruptionfilter.html#cfn-transcribe-callanalyticscategory-interruptionfilter-absolutetimerange
	//
	AbsoluteTimeRange interface{} `field:"optional" json:"absoluteTimeRange" yaml:"absoluteTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-interruptionfilter.html#cfn-transcribe-callanalyticscategory-interruptionfilter-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-interruptionfilter.html#cfn-transcribe-callanalyticscategory-interruptionfilter-participantrole
	//
	ParticipantRole *string `field:"optional" json:"participantRole" yaml:"participantRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-interruptionfilter.html#cfn-transcribe-callanalyticscategory-interruptionfilter-relativetimerange
	//
	RelativeTimeRange interface{} `field:"optional" json:"relativeTimeRange" yaml:"relativeTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-interruptionfilter.html#cfn-transcribe-callanalyticscategory-interruptionfilter-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

