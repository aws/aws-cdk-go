package awsimagebuilder


// Defines filters that the lifecycle policy uses to determine impacted resource.
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
	// Filter resources based on either `age` or `count` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The number of units for the time period or for the count.
	//
	// For example, a value of `6` might refer to six months or six AMIs.
	//
	// > For count-based filters, this value represents the minimum number of resources to keep on hand. If you have fewer resources than this number, the resource is excluded from lifecycle actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// For age-based filters, this is the number of resources to keep on hand after the lifecycle `DELETE` action is applied.
	//
	// Impacted resources are only deleted if you have more than this number of resources. If you have fewer resources than this number, the impacted resource is not deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-retainatleast
	//
	RetainAtLeast *float64 `field:"optional" json:"retainAtLeast" yaml:"retainAtLeast"`
	// Defines the unit of time that the lifecycle policy uses to determine impacted resources.
	//
	// This is required for age-based rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-filter.html#cfn-imagebuilder-lifecyclepolicy-filter-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

