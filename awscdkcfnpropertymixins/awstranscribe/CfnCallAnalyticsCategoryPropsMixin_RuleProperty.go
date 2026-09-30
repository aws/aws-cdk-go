package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   ruleProperty := &RuleProperty{
//   	InterruptionFilter: &InterruptionFilterProperty{
//   		AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   			EndTime: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartTime: jsii.Number(123),
//   		},
//   		Negate: jsii.Boolean(false),
//   		ParticipantRole: jsii.String("participantRole"),
//   		RelativeTimeRange: &RelativeTimeRangeProperty{
//   			EndPercentage: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartPercentage: jsii.Number(123),
//   		},
//   		Threshold: jsii.Number(123),
//   	},
//   	NonTalkTimeFilter: &NonTalkTimeFilterProperty{
//   		AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   			EndTime: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartTime: jsii.Number(123),
//   		},
//   		Negate: jsii.Boolean(false),
//   		RelativeTimeRange: &RelativeTimeRangeProperty{
//   			EndPercentage: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartPercentage: jsii.Number(123),
//   		},
//   		Threshold: jsii.Number(123),
//   	},
//   	SentimentFilter: &SentimentFilterProperty{
//   		AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   			EndTime: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartTime: jsii.Number(123),
//   		},
//   		Negate: jsii.Boolean(false),
//   		ParticipantRole: jsii.String("participantRole"),
//   		RelativeTimeRange: &RelativeTimeRangeProperty{
//   			EndPercentage: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartPercentage: jsii.Number(123),
//   		},
//   		Sentiments: []*string{
//   			jsii.String("sentiments"),
//   		},
//   	},
//   	TranscriptFilter: &TranscriptFilterProperty{
//   		AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   			EndTime: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartTime: jsii.Number(123),
//   		},
//   		Negate: jsii.Boolean(false),
//   		ParticipantRole: jsii.String("participantRole"),
//   		RelativeTimeRange: &RelativeTimeRangeProperty{
//   			EndPercentage: jsii.Number(123),
//   			First: jsii.Number(123),
//   			Last: jsii.Number(123),
//   			StartPercentage: jsii.Number(123),
//   		},
//   		Targets: []*string{
//   			jsii.String("targets"),
//   		},
//   		TranscriptFilterType: jsii.String("transcriptFilterType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-rule.html
//
type CfnCallAnalyticsCategoryPropsMixin_RuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-rule.html#cfn-transcribe-callanalyticscategory-rule-interruptionfilter
	//
	InterruptionFilter interface{} `field:"optional" json:"interruptionFilter" yaml:"interruptionFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-rule.html#cfn-transcribe-callanalyticscategory-rule-nontalktimefilter
	//
	NonTalkTimeFilter interface{} `field:"optional" json:"nonTalkTimeFilter" yaml:"nonTalkTimeFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-rule.html#cfn-transcribe-callanalyticscategory-rule-sentimentfilter
	//
	SentimentFilter interface{} `field:"optional" json:"sentimentFilter" yaml:"sentimentFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-rule.html#cfn-transcribe-callanalyticscategory-rule-transcriptfilter
	//
	TranscriptFilter interface{} `field:"optional" json:"transcriptFilter" yaml:"transcriptFilter"`
}

