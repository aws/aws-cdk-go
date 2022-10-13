// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Properties for an AppSync GraphQL API.
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
type GraphqlApiProps struct {
	// the name of the GraphQL API.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional authorization configuration.
	// Experimental.
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The domain name configuration for the GraphQL API.
	//
	// The Route 53 hosted zone and CName DNS record must be configured in addition to this setting to
	// enable custom domain URL.
	// Experimental.
	DomainName *DomainOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Logging configuration for this api.
	// Experimental.
	LogConfig *LogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// GraphQL schema definition. Specify how you want to define your schema.
	//
	// Schema.fromFile(filePath: string) allows schema definition through schema.graphql file
	// Experimental.
	Schema Schema `field:"optional" json:"schema" yaml:"schema"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	// Experimental.
	XrayEnabled *bool `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

