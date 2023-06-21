package awsappsync


// Configuration of the API authorization modes.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var authFunction function
//
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("appsync.test.graphql"))),
//   	AuthorizationConfig: &AuthorizationConfig{
//   		DefaultAuthorization: &AuthorizationMode{
//   			AuthorizationType: appsync.AuthorizationType_LAMBDA,
//   			LambdaAuthorizerConfig: &LambdaAuthorizerConfig{
//   				Handler: authFunction,
//   			},
//   		},
//   	},
//   })
//
type AuthorizationConfig struct {
	// Additional authorization modes.
	AdditionalAuthorizationModes *[]*AuthorizationMode `field:"optional" json:"additionalAuthorizationModes" yaml:"additionalAuthorizationModes"`
	// Optional authorization configuration.
	DefaultAuthorization *AuthorizationMode `field:"optional" json:"defaultAuthorization" yaml:"defaultAuthorization"`
}

