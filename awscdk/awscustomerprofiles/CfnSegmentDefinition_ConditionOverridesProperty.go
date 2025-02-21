package awscustomerprofiles


// An object to override the original condition block of a calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionOverridesProperty := &ConditionOverridesProperty{
//   	Range: &RangeOverrideProperty{
//   		Start: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//
//   		// the properties below are optional
//   		End: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-conditionoverrides.html
//
type CfnSegmentDefinition_ConditionOverridesProperty struct {
	// The relative time period over which data is included in the aggregation for this override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-conditionoverrides.html#cfn-customerprofiles-segmentdefinition-conditionoverrides-range
	//
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

