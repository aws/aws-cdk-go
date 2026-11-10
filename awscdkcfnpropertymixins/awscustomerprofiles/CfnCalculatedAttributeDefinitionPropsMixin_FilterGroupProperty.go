package awscustomerprofiles


// Object that holds the dimensions to filter on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   filterGroupProperty := &FilterGroupProperty{
//   	Dimensions: []interface{}{
//   		&FilterDimensionProperty{
//   			Attributes: map[string]interface{}{
//   				"attributesKey": &FilterAttributeDimensionProperty{
//   					"dimensionType": jsii.String("dimensionType"),
//   					"values": []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filtergroup.html
//
type CfnCalculatedAttributeDefinitionPropsMixin_FilterGroupProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filtergroup.html#cfn-customerprofiles-calculatedattributedefinition-filtergroup-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The type of logical relationship between the dimensions of the Filter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filtergroup.html#cfn-customerprofiles-calculatedattributedefinition-filtergroup-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

