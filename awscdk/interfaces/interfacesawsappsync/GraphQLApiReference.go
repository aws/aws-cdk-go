package interfacesawsappsync


// A reference to a GraphQLApi resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   graphQLApiReference := &GraphQLApiReference{
//   	ApiId: jsii.String("apiId"),
//   	GraphQlApiArn: jsii.String("graphQlApiArn"),
//   }
//
type GraphQLApiReference struct {
	// The ApiId of the GraphQLApi resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The ARN of the GraphQLApi resource.
	GraphQlApiArn *string `field:"required" json:"graphQlApiArn" yaml:"graphQlApiArn"`
}

