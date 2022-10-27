// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Properties for configuring an Union Type.
//
// Example:
//   var api graphqlApi
//
//   string := appsync.graphqlType.string()
//   human := appsync.NewObjectType(jsii.String("Human"), &objectTypeOptions{
//   	definition: map[string]iField{
//   		"name": string,
//   	},
//   })
//   droid := appsync.NewObjectType(jsii.String("Droid"), &objectTypeOptions{
//   	definition: map[string]*iField{
//   		"name": string,
//   	},
//   })
//   starship := appsync.NewObjectType(jsii.String("Starship"), &objectTypeOptions{
//   	definition: map[string]*iField{
//   		"name": string,
//   	},
//   })
//   search := appsync.NewUnionType(jsii.String("Search"), &unionTypeOptions{
//   	definition: []iIntermediateType{
//   		human,
//   		droid,
//   		starship,
//   	},
//   })
//   api.addType(search)
//
// Experimental.
type UnionTypeOptions struct {
	// the object types for this union type.
	// Experimental.
	Definition *[]IIntermediateType `field:"required" json:"definition" yaml:"definition"`
}

