package awscustomerprofiles


// Object that defines how to filter the incoming objects for the calculated attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filterAttributeDimensionProperty := &FilterAttributeDimensionProperty{
//   	DimensionType: jsii.String("dimensionType"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filterattributedimension.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_FilterAttributeDimensionProperty struct {
	// The action to filter with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filterattributedimension.html#cfn-customerprofiles-calculatedattributedefinition-filterattributedimension-dimensiontype
	//
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The values to apply the DimensionType on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filterattributedimension.html#cfn-customerprofiles-calculatedattributedefinition-filterattributedimension-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

