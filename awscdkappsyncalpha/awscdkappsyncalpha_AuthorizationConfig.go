// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Configuration of the API authorization modes.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var authFunction function
//
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schema.fromAsset(path.join(__dirname, jsii.String("appsync.test.graphql"))),
//   	authorizationConfig: &authorizationConfig{
//   		defaultAuthorization: &authorizationMode{
//   			authorizationType: appsync.authorizationType_LAMBDA,
//   			lambdaAuthorizerConfig: &lambdaAuthorizerConfig{
//   				handler: authFunction,
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type AuthorizationConfig struct {
	// Additional authorization modes.
	// Experimental.
	AdditionalAuthorizationModes *[]*AuthorizationMode `field:"optional" json:"additionalAuthorizationModes" yaml:"additionalAuthorizationModes"`
	// Optional authorization configuration.
	// Experimental.
	DefaultAuthorization *AuthorizationMode `field:"optional" json:"defaultAuthorization" yaml:"defaultAuthorization"`
}

