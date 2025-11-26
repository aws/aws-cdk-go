package previewawscustomerprofilesmixins


// Object that segments on various Customer profile's fields that are larger than normal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   profileDimensionProperty := &ProfileDimensionProperty{
//   	DimensionType: jsii.String("dimensionType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiledimension.html
//
type CfnSegmentDefinitionPropsMixin_ProfileDimensionProperty struct {
	// The action to segment on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiledimension.html#cfn-customerprofiles-segmentdefinition-profiledimension-dimensiontype
	//
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-profiledimension.html#cfn-customerprofiles-segmentdefinition-profiledimension-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

