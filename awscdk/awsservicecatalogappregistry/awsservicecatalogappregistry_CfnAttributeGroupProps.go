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
//   cfnAttributeGroupProps := &cfnAttributeGroupProps{
//   	attributes: attributes,
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnAttributeGroupProps struct {
	// A JSON string in the form of nested key-value pairs that represent the attributes in the group and describes an application and its components.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the attribute group.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the attribute group that the user provides.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs you can use to associate with the attribute group.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

