package awsappsync


// Introspection configuration  for a GraphQL API.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("DisableIntrospectionApi"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
//   	IntrospectionConfig: appsync.IntrospectionConfig_DISABLED,
//   })
//
type IntrospectionConfig string

const (
	// Enable introspection.
	IntrospectionConfig_ENABLED IntrospectionConfig = "ENABLED"
	// Disable introspection.
	IntrospectionConfig_DISABLED IntrospectionConfig = "DISABLED"
)

