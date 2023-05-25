package awsappsync


// Attributes for GraphQL imports.
//
// Example:
//   var api graphqlApi
//   var table table
//
//   importedApi := appsync.graphqlApi_FromGraphqlApiAttributes(this, jsii.String("IApi"), &GraphqlApiAttributes{
//   	GraphqlApiId: api.ApiId,
//   	GraphqlApiArn: api.Arn,
//   })
//   importedApi.AddDynamoDbDataSource(jsii.String("TableDataSource"), table)
//
type GraphqlApiAttributes struct {
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	GraphqlApiId *string `field:"required" json:"graphqlApiId" yaml:"graphqlApiId"`
	// the arn for the GraphQL Api.
	GraphqlApiArn *string `field:"optional" json:"graphqlApiArn" yaml:"graphqlApiArn"`
}

