package awsinspectorv2


// An object that describes the details of a string filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringFilterProperty := &StringFilterProperty{
//   	Comparison: jsii.String("comparison"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-stringfilter.html
//
type CfnFilter_StringFilterProperty struct {
	// The operator to use when comparing values in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-stringfilter.html#cfn-inspectorv2-filter-stringfilter-comparison
	//
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The value to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-stringfilter.html#cfn-inspectorv2-filter-stringfilter-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

