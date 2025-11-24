package mixinsawscustomerprofiles


// Overrides the original range on a calculated attribute definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rangeOverrideProperty := &RangeOverrideProperty{
//   	End: jsii.Number(123),
//   	Start: jsii.Number(123),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-rangeoverride.html
//
type CfnSegmentDefinitionPropsMixin_RangeOverrideProperty struct {
	// The end time of when to include objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-rangeoverride.html#cfn-customerprofiles-segmentdefinition-rangeoverride-end
	//
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// The start time of when to include objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-rangeoverride.html#cfn-customerprofiles-segmentdefinition-rangeoverride-start
	//
	Start *float64 `field:"optional" json:"start" yaml:"start"`
	// The unit for start and end.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-rangeoverride.html#cfn-customerprofiles-segmentdefinition-rangeoverride-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

