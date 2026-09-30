package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCallAnalyticsCategory`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCallAnalyticsCategoryProps := &CfnCallAnalyticsCategoryProps{
//   	CategoryName: jsii.String("categoryName"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			InterruptionFilter: &InterruptionFilterProperty{
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				ParticipantRole: jsii.String("participantRole"),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   				Threshold: jsii.Number(123),
//   			},
//   			NonTalkTimeFilter: &NonTalkTimeFilterProperty{
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   				Threshold: jsii.Number(123),
//   			},
//   			SentimentFilter: &SentimentFilterProperty{
//   				Sentiments: []*string{
//   					jsii.String("sentiments"),
//   				},
//
//   				// the properties below are optional
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				ParticipantRole: jsii.String("participantRole"),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   			},
//   			TranscriptFilter: &TranscriptFilterProperty{
//   				Targets: []*string{
//   					jsii.String("targets"),
//   				},
//   				TranscriptFilterType: jsii.String("transcriptFilterType"),
//
//   				// the properties below are optional
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				ParticipantRole: jsii.String("participantRole"),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	InputType: jsii.String("inputType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html
//
type CfnCallAnalyticsCategoryProps struct {
	// A unique name, chosen by you, for your Call Analytics category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-categoryname
	//
	CategoryName *string `field:"required" json:"categoryName" yaml:"categoryName"`
	// Rules define a Call Analytics category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The input type associated with the specified category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-inputtype
	//
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// Tags associated with the Call Analytics category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

