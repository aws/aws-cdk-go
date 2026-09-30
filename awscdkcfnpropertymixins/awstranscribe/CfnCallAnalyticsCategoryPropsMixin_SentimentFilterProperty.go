package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sentimentFilterProperty := &SentimentFilterProperty{
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
//   	Sentiments: []*string{
//   		jsii.String("sentiments"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-sentimentfilter.html
//
type CfnCallAnalyticsCategoryPropsMixin_SentimentFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-sentimentfilter.html#cfn-transcribe-callanalyticscategory-sentimentfilter-absolutetimerange
	//
	AbsoluteTimeRange interface{} `field:"optional" json:"absoluteTimeRange" yaml:"absoluteTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-sentimentfilter.html#cfn-transcribe-callanalyticscategory-sentimentfilter-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-sentimentfilter.html#cfn-transcribe-callanalyticscategory-sentimentfilter-participantrole
	//
	ParticipantRole *string `field:"optional" json:"participantRole" yaml:"participantRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-sentimentfilter.html#cfn-transcribe-callanalyticscategory-sentimentfilter-relativetimerange
	//
	RelativeTimeRange interface{} `field:"optional" json:"relativeTimeRange" yaml:"relativeTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-sentimentfilter.html#cfn-transcribe-callanalyticscategory-sentimentfilter-sentiments
	//
	Sentiments *[]*string `field:"optional" json:"sentiments" yaml:"sentiments"`
}

