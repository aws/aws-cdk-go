package awscustomerprofiles


// Defines how to filter the objects coming in for calculated attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filterProperty := &FilterProperty{
//   	Groups: []interface{}{
//   		&FilterGroupProperty{
//   			Dimensions: []interface{}{
//   				&FilterDimensionProperty{
//   					Attributes: map[string]interface{}{
//   						"attributesKey": &FilterAttributeDimensionProperty{
//   							"dimensionType": jsii.String("dimensionType"),
//   							"values": []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Include: jsii.String("include"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filter.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_FilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filter.html#cfn-customerprofiles-calculatedattributedefinition-filter-groups
	//
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// Define whether to include or exclude objects for Calculated Attributed calculation that fit the filter groups criteria.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filter.html#cfn-customerprofiles-calculatedattributedefinition-filter-include
	//
	Include *string `field:"optional" json:"include" yaml:"include"`
}

