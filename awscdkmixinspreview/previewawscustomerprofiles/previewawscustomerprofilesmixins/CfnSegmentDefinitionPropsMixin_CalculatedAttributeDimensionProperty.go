package previewawscustomerprofilesmixins


// Object that segments on Customer Profile's Calculated Attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   calculatedAttributeDimensionProperty := &CalculatedAttributeDimensionProperty{
//   	ConditionOverrides: &ConditionOverridesProperty{
//   		Range: &RangeOverrideProperty{
//   			End: jsii.Number(123),
//   			Start: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   	DimensionType: jsii.String("dimensionType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html
//
type CfnSegmentDefinitionPropsMixin_CalculatedAttributeDimensionProperty struct {
	// Applies the given condition over the initial Calculated Attribute's definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html#cfn-customerprofiles-segmentdefinition-calculatedattributedimension-conditionoverrides
	//
	ConditionOverrides interface{} `field:"optional" json:"conditionOverrides" yaml:"conditionOverrides"`
	// The action to segment with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html#cfn-customerprofiles-segmentdefinition-calculatedattributedimension-dimensiontype
	//
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The values to apply the DimensionType with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-calculatedattributedimension.html#cfn-customerprofiles-segmentdefinition-calculatedattributedimension-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

