package awscustomerprofiles


// Defines the characteristics and rules for sorting by a specific attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   sortAttributeProperty := &SortAttributeProperty{
//   	DataType: jsii.String("dataType"),
//   	Name: jsii.String("name"),
//   	Order: jsii.String("order"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sortattribute.html
//
type CfnSegmentDefinitionPropsMixin_SortAttributeProperty struct {
	// The data type of the sort attribute (e.g., string, number, date).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sortattribute.html#cfn-customerprofiles-segmentdefinition-sortattribute-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The name of the attribute to sort by.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sortattribute.html#cfn-customerprofiles-segmentdefinition-sortattribute-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order for the attribute (ascending or descending).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sortattribute.html#cfn-customerprofiles-segmentdefinition-sortattribute-order
	//
	Order *string `field:"optional" json:"order" yaml:"order"`
	// The type of attribute (e.g., profile, calculated).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-segmentdefinition-sortattribute.html#cfn-customerprofiles-segmentdefinition-sortattribute-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

