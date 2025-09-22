package awsappsync


// A reference to a GraphQLSchema resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   graphQLSchemaReference := &GraphQLSchemaReference{
//   	GraphQlSchemaId: jsii.String("graphQlSchemaId"),
//   }
//
type GraphQLSchemaReference struct {
	// The Id of the GraphQLSchema resource.
	GraphQlSchemaId *string `field:"required" json:"graphQlSchemaId" yaml:"graphQlSchemaId"`
}

