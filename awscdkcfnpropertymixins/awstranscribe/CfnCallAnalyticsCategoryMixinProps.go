package awstranscribe

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCallAnalyticsCategoryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCallAnalyticsCategoryMixinProps := &CfnCallAnalyticsCategoryMixinProps{
//   	CategoryName: jsii.String("categoryName"),
//   	InputType: jsii.String("inputType"),
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
//   				Sentiments: []*string{
//   					jsii.String("sentiments"),
//   				},
//   			},
//   			TranscriptFilter: &TranscriptFilterProperty{
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
//   				Targets: []*string{
//   					jsii.String("targets"),
//   				},
//   				TranscriptFilterType: jsii.String("transcriptFilterType"),
//   			},
//   		},
//   	},
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
type CfnCallAnalyticsCategoryMixinProps struct {
	// A unique name, chosen by you, for your Call Analytics category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-categoryname
	//
	CategoryName *string `field:"optional" json:"categoryName" yaml:"categoryName"`
	// The input type associated with the specified category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-inputtype
	//
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// Rules define a Call Analytics category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Tags associated with the Call Analytics category.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html#cfn-transcribe-callanalyticscategory-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

