package awsinspectorv2


// An object that describes details of a map filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapFilterProperty := &MapFilterProperty{
//   	Comparison: jsii.String("comparison"),
//
//   	// the properties below are optional
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-mapfilter.html
//
type CfnFilter_MapFilterProperty struct {
	// The operator to use when comparing values in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-mapfilter.html#cfn-inspectorv2-filter-mapfilter-comparison
	//
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// The tag key used in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-mapfilter.html#cfn-inspectorv2-filter-mapfilter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value used in the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-filter-mapfilter.html#cfn-inspectorv2-filter-mapfilter-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

