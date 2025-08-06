package awscustomerprofiles


// Object that segments on Customer Profile's Calculated Attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   calculatedAttributeDimensionProperty := &CalculatedAttributeDimensionProperty{
//   	DimensionType: jsii.String("dimensionType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//
//   	// the properties below are optional
//   	ConditionOverrides: &ConditionOverridesProperty{
//   		Range: &RangeOverrideProperty{
//   			Start: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//
//   			// the properties below are optional
//   			End: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html
//
type CfnSegmentDefinition_CalculatedAttributeDimensionProperty struct {
	// The action to segment with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html#cfn-customerprofiles-segmentdefinition-calculatedattributedimension-dimensiontype
	//
	DimensionType *string `field:"required" json:"dimensionType" yaml:"dimensionType"`
	// The values to apply the DimensionType with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html#cfn-customerprofiles-segmentdefinition-calculatedattributedimension-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Applies the given condition over the initial Calculated Attribute's definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html#cfn-customerprofiles-segmentdefinition-calculatedattributedimension-conditionoverrides
	//
	ConditionOverrides interface{} `field:"optional" json:"conditionOverrides" yaml:"conditionOverrides"`
}

