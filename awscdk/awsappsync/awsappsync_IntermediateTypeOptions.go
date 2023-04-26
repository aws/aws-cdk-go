package awsappsync


// Properties for configuring an Intermediate Type.
//
// Example:
//   node := appsync.NewInterfaceType(jsii.String("Node"), &IntermediateTypeOptions{
//   	Definition: map[string]iField{
//   		"id": appsync.GraphqlType_string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   })
//   demo := appsync.NewObjectType(jsii.String("Demo"), &ObjectTypeOptions{
//   	InterfaceTypes: []interfaceType{
//   		node,
//   	},
//   	Definition: map[string]*iField{
//   		"version": appsync.GraphqlType_string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   })
//
// Experimental.
type IntermediateTypeOptions struct {
	// the attributes of this type.
	// Experimental.
	Definition *map[string]IField `field:"required" json:"definition" yaml:"definition"`
	// the directives for this object type.
	// Experimental.
	Directives *[]Directive `field:"optional" json:"directives" yaml:"directives"`
}

