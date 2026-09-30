package awstranscribe


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   relativeTimeRangeProperty := &RelativeTimeRangeProperty{
//   	EndPercentage: jsii.Number(123),
//   	First: jsii.Number(123),
//   	Last: jsii.Number(123),
//   	StartPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-relativetimerange.html
//
type CfnCallAnalyticsCategoryPropsMixin_RelativeTimeRangeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-relativetimerange.html#cfn-transcribe-callanalyticscategory-relativetimerange-endpercentage
	//
	EndPercentage *float64 `field:"optional" json:"endPercentage" yaml:"endPercentage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-relativetimerange.html#cfn-transcribe-callanalyticscategory-relativetimerange-first
	//
	First *float64 `field:"optional" json:"first" yaml:"first"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-relativetimerange.html#cfn-transcribe-callanalyticscategory-relativetimerange-last
	//
	Last *float64 `field:"optional" json:"last" yaml:"last"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transcribe-callanalyticscategory-relativetimerange.html#cfn-transcribe-callanalyticscategory-relativetimerange-startpercentage
	//
	StartPercentage *float64 `field:"optional" json:"startPercentage" yaml:"startPercentage"`
}

