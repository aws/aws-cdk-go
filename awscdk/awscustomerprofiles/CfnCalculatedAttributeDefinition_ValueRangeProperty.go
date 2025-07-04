package awscustomerprofiles


// A structure specifying the endpoints of the relative time period over which data is included in the aggregation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueRangeProperty := &ValueRangeProperty{
//   	End: jsii.Number(123),
//   	Start: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-valuerange.html
//
type CfnCalculatedAttributeDefinition_ValueRangeProperty struct {
	// The ending point for this range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-valuerange.html#cfn-customerprofiles-calculatedattributedefinition-valuerange-end
	//
	End *float64 `field:"required" json:"end" yaml:"end"`
	// The starting point for this range.
	//
	// Positive numbers indicate how many days in the past data should be included, and negative numbers indicate how many days in the future.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-valuerange.html#cfn-customerprofiles-calculatedattributedefinition-valuerange-start
	//
	Start *float64 `field:"required" json:"start" yaml:"start"`
}

