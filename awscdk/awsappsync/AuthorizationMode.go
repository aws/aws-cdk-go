package awsappsync


// Interface to specify default or additional authorization(s).
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Definition: appsync.Definition_FromFile(jsii.String("schema.graphql")),
//   	AuthorizationConfig: &AuthorizationConfig{
//   		DefaultAuthorization: &AuthorizationMode{
//   			AuthorizationType: appsync.AuthorizationType_IAM,
//   		},
//   	},
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewAppSync(api, &AppSyncGraphQLApiProps{
//   	GraphQLOperation: jsii.String("mutation Publish($message: String!){ publish(message: $message) { message } }"),
//   	Variables: events.RuleTargetInput_FromObject(map[string]*string{
//   		"message": jsii.String("hello world"),
//   	}),
//   }))
//
type AuthorizationMode struct {
	// One of possible four values AppSync supports.
	// See: https://docs.aws.amazon.com/appsync/latest/devguide/security.html
	//
	// Default: - `AuthorizationType.API_KEY`
	//
	AuthorizationType AuthorizationType `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// If authorizationType is `AuthorizationType.API_KEY`, this option can be configured.
	// Default: - name: 'DefaultAPIKey' | description: 'Default API Key created by CDK'.
	//
	ApiKeyConfig *ApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// If authorizationType is `AuthorizationType.LAMBDA`, this option is required.
	// Default: - none.
	//
	LambdaAuthorizerConfig *LambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// If authorizationType is `AuthorizationType.OIDC`, this option is required.
	// Default: - none.
	//
	OpenIdConnectConfig *OpenIdConnectConfig `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
	// If authorizationType is `AuthorizationType.USER_POOL`, this option is required.
	// Default: - none.
	//
	UserPoolConfig *UserPoolConfig `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
}

