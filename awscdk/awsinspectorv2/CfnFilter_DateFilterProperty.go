package awsinspectorv2


// Contains details on the time range used to filter findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateFilterProperty := &DateFilterProperty{
//   	EndInclusive: jsii.Number(123),
//   	StartInclusive: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-datefilter.html
//
type CfnFilter_DateFilterProperty struct {
	// A timestamp representing the end of the time period filtered on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-datefilter.html#cfn-inspectorv2-filter-datefilter-endinclusive
	//
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
	// A timestamp representing the start of the time period filtered on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-datefilter.html#cfn-inspectorv2-filter-datefilter-startinclusive
	//
	StartInclusive *float64 `field:"optional" json:"startInclusive" yaml:"startInclusive"`
}

