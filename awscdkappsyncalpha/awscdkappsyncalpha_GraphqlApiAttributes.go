// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Attributes for GraphQL imports.
//
// Example:
//   var api graphqlApi
//   var table table
//
//   importedApi := appsync.graphqlApi.fromGraphqlApiAttributes(this, jsii.String("IApi"), &graphqlApiAttributes{
//   	graphqlApiId: api.apiId,
//   	graphqlApiArn: api.arn,
//   })
//   importedApi.addDynamoDbDataSource(jsii.String("TableDataSource"), table)
//
// Experimental.
type GraphqlApiAttributes struct {
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	// Experimental.
	GraphqlApiId *string `field:"required" json:"graphqlApiId" yaml:"graphqlApiId"`
	// the arn for the GraphQL Api.
	// Experimental.
	GraphqlApiArn *string `field:"optional" json:"graphqlApiArn" yaml:"graphqlApiArn"`
}

