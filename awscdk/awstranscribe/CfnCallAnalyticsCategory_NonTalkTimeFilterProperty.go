package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nonTalkTimeFilterProperty := &NonTalkTimeFilterProperty{
//   	AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   		EndTime: jsii.Number(123),
//   		First: jsii.Number(123),
//   		Last: jsii.Number(123),
//   		StartTime: jsii.Number(123),
//   	},
//   	Negate: jsii.Boolean(false),
//   	RelativeTimeRange: &RelativeTimeRangeProperty{
//   		EndPercentage: jsii.Number(123),
//   		First: jsii.Number(123),
//   		Last: jsii.Number(123),
//   		StartPercentage: jsii.Number(123),
//   	},
//   	Threshold: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-nontalktimefilter.html
//
type CfnCallAnalyticsCategory_NonTalkTimeFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-nontalktimefilter.html#cfn-transcribe-callanalyticscategory-nontalktimefilter-absolutetimerange
	//
	AbsoluteTimeRange interface{} `field:"optional" json:"absoluteTimeRange" yaml:"absoluteTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-nontalktimefilter.html#cfn-transcribe-callanalyticscategory-nontalktimefilter-negate
	//
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-nontalktimefilter.html#cfn-transcribe-callanalyticscategory-nontalktimefilter-relativetimerange
	//
	RelativeTimeRange interface{} `field:"optional" json:"relativeTimeRange" yaml:"relativeTimeRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-nontalktimefilter.html#cfn-transcribe-callanalyticscategory-nontalktimefilter-threshold
	//
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

