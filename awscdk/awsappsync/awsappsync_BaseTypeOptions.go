package awsappsync


// Base options for GraphQL Types.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("Api"), &graphqlApiProps{
//   	name: jsii.String("demo"),
//   })
//   demo := appsync.NewObjectType(jsii.String("Demo"), &objectTypeOptions{
//   	definition: map[string]iField{
//   		"id": appsync.GraphqlType.string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   		"version": appsync.GraphqlType.string(&BaseTypeOptions{
//   			"isRequired": jsii.Boolean(true),
//   		}),
//   	},
//   })
//
//   api.addType(demo)
//
// Experimental.
type BaseTypeOptions struct {
	// property determining if this attribute is a list i.e. if true, attribute would be [Type].
	// Experimental.
	IsList *bool `field:"optional" json:"isList" yaml:"isList"`
	// property determining if this attribute is non-nullable i.e. if true, attribute would be Type!
	// Experimental.
	IsRequired *bool `field:"optional" json:"isRequired" yaml:"isRequired"`
	// property determining if this attribute is a non-nullable list i.e. if true, attribute would be [ Type ]! or if isRequired true, attribe would be [ Type! ]!
	// Experimental.
	IsRequiredList *bool `field:"optional" json:"isRequiredList" yaml:"isRequiredList"`
}

