package awsservicecatalogappregistry


// Properties for a Service Catalog AppRegistry Attribute Group.
//
// Example:
//   attributeGroup := appreg.NewAttributeGroup(this, jsii.String("MyFirstAttributeGroup"), &AttributeGroupProps{
//   	AttributeGroupName: jsii.String("MyFirstAttributeGroupName"),
//   	Description: jsii.String("description for my attribute group"),
//   	 // the description is optional,
//   	Attributes: map[string]interface{}{
//   		"project": jsii.String("foo"),
//   		"team": []interface{}{
//   			jsii.String("member1"),
//   			jsii.String("member2"),
//   			jsii.String("member3"),
//   		},
//   		"public": jsii.Boolean(false),
//   		"stages": map[string]*string{
//   			"alpha": jsii.String("complete"),
//   			"beta": jsii.String("incomplete"),
//   			"release": jsii.String("not started"),
//   		},
//   	},
//   })
//
// Experimental.
type AttributeGroupProps struct {
	// Enforces a particular physical attribute group name.
	// Experimental.
	AttributeGroupName *string `field:"required" json:"attributeGroupName" yaml:"attributeGroupName"`
	// A JSON of nested key-value pairs that represent the attributes in the group.
	//
	// Attributes maybe an empty JSON '{}', but must be explicitly stated.
	// Experimental.
	Attributes *map[string]interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// Description for attribute group.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

