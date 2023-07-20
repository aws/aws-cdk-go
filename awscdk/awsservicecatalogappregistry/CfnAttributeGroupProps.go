package awsservicecatalogappregistry


// Properties for defining a `CfnAttributeGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//
//   cfnAttributeGroupProps := &CfnAttributeGroupProps{
//   	Attributes: attributes,
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroup.html
//
type CfnAttributeGroupProps struct {
	// A nested object in a JSON or YAML template that supports arbitrary definitions.
	//
	// Represents the attributes in an attribute group that describes an application and its components.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroup.html#cfn-servicecatalogappregistry-attributegroup-attributes
	//
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the attribute group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroup.html#cfn-servicecatalogappregistry-attributegroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the attribute group that the user provides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroup.html#cfn-servicecatalogappregistry-attributegroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs you can use to associate with the attribute group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroup.html#cfn-servicecatalogappregistry-attributegroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

