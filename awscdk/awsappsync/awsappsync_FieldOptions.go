package awsappsync


// Properties for configuring a field.
//
// Example:
//   field := appsync.NewField(&fieldOptions{
//   	returnType: appsync.graphqlType.string(),
//   	args: map[string]*graphqlType{
//   		"argument": appsync.*graphqlType.string(),
//   	},
//   })
//   type := appsync.NewInterfaceType(jsii.String("Node"), &intermediateTypeOptions{
//   	definition: map[string]iField{
//   		"test": field,
//   	},
//   })
//
// Experimental.
type FieldOptions struct {
	// The return type for this field.
	// Experimental.
	ReturnType GraphqlType `field:"required" json:"returnType" yaml:"returnType"`
	// The arguments for this field.
	//
	// i.e. type Example (first: String second: String) {}
	// - where 'first' and 'second' are key values for args
	// and 'String' is the GraphqlType.
	// Experimental.
	Args *map[string]GraphqlType `field:"optional" json:"args" yaml:"args"`
	// the directives for this field.
	// Experimental.
	Directives *[]Directive `field:"optional" json:"directives" yaml:"directives"`
}

