package previewawsinspectorv2mixins


// An object that describes the details of a number filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   numberFilterProperty := &NumberFilterProperty{
//   	LowerInclusive: jsii.Number(123),
//   	UpperInclusive: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-numberfilter.html
//
type CfnFilterPropsMixin_NumberFilterProperty struct {
	// The lowest number to be included in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-numberfilter.html#cfn-inspectorv2-filter-numberfilter-lowerinclusive
	//
	LowerInclusive *float64 `field:"optional" json:"lowerInclusive" yaml:"lowerInclusive"`
	// The highest number to be included in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-numberfilter.html#cfn-inspectorv2-filter-numberfilter-upperinclusive
	//
	UpperInclusive *float64 `field:"optional" json:"upperInclusive" yaml:"upperInclusive"`
}

