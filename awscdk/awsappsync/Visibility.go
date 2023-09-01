package awsappsync


// Visibility type for a GraphQL API.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("MyPrivateAPI"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.schema.graphql"))),
//   	Visibility: appsync.Visibility_PRIVATE,
//   })
//
type Visibility string

const (
	// Public, open to the internet.
	Visibility_GLOBAL Visibility = "GLOBAL"
	// Only accessible through a VPC.
	Visibility_PRIVATE Visibility = "PRIVATE"
)

