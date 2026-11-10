package awscustomerprofiles


// Contains the map of attribute names to attribute dimensions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterDimensionProperty := &FilterDimensionProperty{
//   	Attributes: map[string]interface{}{
//   		"attributesKey": &FilterAttributeDimensionProperty{
//   			"dimensionType": jsii.String("dimensionType"),
//   			"values": []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filterdimension.html
//
type CfnCalculatedAttributeDefinition_FilterDimensionProperty struct {
	// Is the attribute within the FilterDimension map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-calculatedattributedefinition-filterdimension.html#cfn-customerprofiles-calculatedattributedefinition-filterdimension-attributes
	//
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
}

