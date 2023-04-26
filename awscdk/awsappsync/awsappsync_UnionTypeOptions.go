package awsappsync


// Properties for configuring an Union Type.
//
// Example:
//   var api graphqlApi
//
//   string := appsync.GraphqlType_String()
//   human := appsync.NewObjectType(jsii.String("Human"), &ObjectTypeOptions{
//   	Definition: map[string]iField{
//   		"name": string,
//   	},
//   })
//   droid := appsync.NewObjectType(jsii.String("Droid"), &ObjectTypeOptions{
//   	Definition: map[string]*iField{
//   		"name": string,
//   	},
//   })
//   starship := appsync.NewObjectType(jsii.String("Starship"), &ObjectTypeOptions{
//   	Definition: map[string]*iField{
//   		"name": string,
//   	},
//   })
//   search := appsync.NewUnionType(jsii.String("Search"), &UnionTypeOptions{
//   	Definition: []iIntermediateType{
//   		human,
//   		droid,
//   		starship,
//   	},
//   })
//   api.AddType(search)
//
// Experimental.
type UnionTypeOptions struct {
	// the object types for this union type.
	// Experimental.
	Definition *[]IIntermediateType `field:"required" json:"definition" yaml:"definition"`
}

