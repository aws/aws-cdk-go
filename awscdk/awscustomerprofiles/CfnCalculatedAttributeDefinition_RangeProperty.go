package awscustomerprofiles


// The relative time period over which data is included in the aggregation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rangeProperty := &RangeProperty{
//   	Unit: jsii.String("unit"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html
//
type CfnCalculatedAttributeDefinition_RangeProperty struct {
	// The unit of time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-unit
	//
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The amount of time of the specified unit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-range.html#cfn-customerprofiles-calculatedattributedefinition-range-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

