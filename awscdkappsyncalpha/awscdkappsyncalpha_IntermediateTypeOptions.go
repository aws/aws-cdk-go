// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Properties for configuring an Intermediate Type.
//
// Example:
//   node := appsync.NewInterfaceType(jsii.String("Node"), &intermediateTypeOptions{
//   	definition: map[string]iField{
//   		"id": appsync.GraphqlType.string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   })
//   demo := appsync.NewObjectType(jsii.String("Demo"), &objectTypeOptions{
//   	interfaceTypes: []interfaceType{
//   		node,
//   	},
//   	definition: map[string]*iField{
//   		"version": appsync.GraphqlType.string(&BaseTypeOptions{
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

