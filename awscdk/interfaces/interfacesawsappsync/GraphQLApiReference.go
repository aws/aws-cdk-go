package interfacesawsappsync


// A reference to a GraphQLApi resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   graphQLApiReference := &GraphQLApiReference{
//   	GraphQlApiArn: jsii.String("graphQlApiArn"),
//   }
//
type GraphQLApiReference struct {
	// The Arn of the GraphQLApi resource.
	GraphQlApiArn *string `field:"required" json:"graphQlApiArn" yaml:"graphQlApiArn"`
}

