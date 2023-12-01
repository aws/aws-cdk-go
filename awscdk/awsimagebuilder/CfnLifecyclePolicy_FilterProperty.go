package awsimagebuilder


// A filter name and value pair that is used to return a more specific list of results from a list operation.
//
// Filters can be used to match a set of resources by specific criteria, such as tags, attributes, or IDs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &FilterProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//
//   	// the properties below are optional
//   	RetainAtLeast: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html
//
type CfnLifecyclePolicy_FilterProperty struct {
	// The filter type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// The minimum number of Image Builder resources to retain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-retainatleast
	//
	RetainAtLeast *float64 `field:"optional" json:"retainAtLeast" yaml:"retainAtLeast"`
	// A time unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

