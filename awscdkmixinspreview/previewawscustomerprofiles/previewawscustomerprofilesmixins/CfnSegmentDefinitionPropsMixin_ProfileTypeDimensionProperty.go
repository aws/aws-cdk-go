package previewawscustomerprofilesmixins


// Specifies profile type based criteria for a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   profileTypeDimensionProperty := &ProfileTypeDimensionProperty{
//   	DimensionType: jsii.String("dimensionType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiletypedimension.html
//
type CfnSegmentDefinitionPropsMixin_ProfileTypeDimensionProperty struct {
	// The action to segment on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiletypedimension.html#cfn-customerprofiles-segmentdefinition-profiletypedimension-dimensiontype
	//
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The values to apply the DimensionType on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiletypedimension.html#cfn-customerprofiles-segmentdefinition-profiletypedimension-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

