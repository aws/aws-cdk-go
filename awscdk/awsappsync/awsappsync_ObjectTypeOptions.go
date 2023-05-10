package awsappsync


// Properties for configuring an Object Type.
//
// Example:
//   var api graphqlApi
//   var dummyRequest mappingTemplate
//   var dummyResponse mappingTemplate
//
//   info := appsync.NewObjectType(jsii.String("Info"), &ObjectTypeOptions{
//   	Definition: map[string]iField{
//   		"node": appsync.NewResolvableField(&ResolvableFieldOptions{
//   			"returnType": appsync.GraphqlType_string(),
//   			"args": map[string]GraphqlType{
//   				"id": appsync.GraphqlType_string(),
//   			},
//   			"dataSource": api.addNoneDataSource(jsii.String("none")),
//   			"requestMappingTemplate": dummyRequest,
//   			"responseMappingTemplate": dummyResponse,
//   		}),
//   	},
//   })
//
// Experimental.
type ObjectTypeOptions struct {
	// the attributes of this type.
	// Experimental.
	Definition *map[string]IField `field:"required" json:"definition" yaml:"definition"`
	// the directives for this object type.
	// Experimental.
	Directives *[]Directive `field:"optional" json:"directives" yaml:"directives"`
	// The Interface Types this Object Type implements.
	// Experimental.
	InterfaceTypes *[]InterfaceType `field:"optional" json:"interfaceTypes" yaml:"interfaceTypes"`
}

