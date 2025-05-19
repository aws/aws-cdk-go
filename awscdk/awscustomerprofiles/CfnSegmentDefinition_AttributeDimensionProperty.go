package awscustomerprofiles


// Object that defines how to filter the incoming objects for the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDimensionProperty := &AttributeDimensionProperty{
//   	DimensionType: jsii.String("dimensionType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-attributedimension.html
//
type CfnSegmentDefinition_AttributeDimensionProperty struct {
	// The action to segment with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-attributedimension.html#cfn-customerprofiles-segmentdefinition-attributedimension-dimensiontype
	//
	DimensionType *string `field:"required" json:"dimensionType" yaml:"dimensionType"`
	// The values to apply the DimensionType on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-attributedimension.html#cfn-customerprofiles-segmentdefinition-attributedimension-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

