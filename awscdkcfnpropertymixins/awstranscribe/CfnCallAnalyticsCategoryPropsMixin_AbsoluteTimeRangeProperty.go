package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   absoluteTimeRangeProperty := &AbsoluteTimeRangeProperty{
//   	EndTime: jsii.Number(123),
//   	First: jsii.Number(123),
//   	Last: jsii.Number(123),
//   	StartTime: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-absolutetimerange.html
//
type CfnCallAnalyticsCategoryPropsMixin_AbsoluteTimeRangeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-absolutetimerange.html#cfn-transcribe-callanalyticscategory-absolutetimerange-endtime
	//
	EndTime *float64 `field:"optional" json:"endTime" yaml:"endTime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-absolutetimerange.html#cfn-transcribe-callanalyticscategory-absolutetimerange-first
	//
	First *float64 `field:"optional" json:"first" yaml:"first"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-absolutetimerange.html#cfn-transcribe-callanalyticscategory-absolutetimerange-last
	//
	Last *float64 `field:"optional" json:"last" yaml:"last"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-absolutetimerange.html#cfn-transcribe-callanalyticscategory-absolutetimerange-starttime
	//
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
}

