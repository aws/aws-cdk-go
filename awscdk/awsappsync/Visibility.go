package awsappsync


// Visibility type for a GraphQL API.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   api := appsync.NewGraphqlApi(stack, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("MyPrivateAPI"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
//   	Visbility: appsync.Visibility_PRIVATE,
//   })
//
type Visibility string

const (
	// Public, open to the internet.
	Visibility_GLOBAL Visibility = "GLOBAL"
	// Only accessible through a VPC.
	Visibility_PRIVATE Visibility = "PRIVATE"
)

