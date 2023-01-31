package awsappsync


// Properties for configuring an Enum Type.
//
// Example:
//   var api graphqlApi
//
//   episode := appsync.NewEnumType(jsii.String("Episode"), &enumTypeOptions{
//   	definition: []*string{
//   		jsii.String("NEWHOPE"),
//   		jsii.String("EMPIRE"),
//   		jsii.String("JEDI"),
//   	},
//   })
//   api.addType(episode)
//
// Experimental.
type EnumTypeOptions struct {
	// the attributes of this type.
	// Experimental.
	Definition *[]*string `field:"required" json:"definition" yaml:"definition"`
}

