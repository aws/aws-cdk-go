package awsappsync


// Auth provider settings for AppSync Event APIs.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   apiKeyProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
//   }
//
//   api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
//   	ApiName: jsii.String("Api"),
//   	OwnerContact: jsii.String("OwnerContact"),
//   	AuthorizationConfig: &EventApiAuthConfig{
//   		AuthProviders: []appSyncAuthProvider{
//   			apiKeyProvider,
//   		},
//   		ConnectionAuthModeTypes: []appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   	},
//   	LogConfig: &AppSyncLogConfig{
//   		FieldLogLevel: appsync.AppSyncFieldLogLevel_INFO,
//   		Retention: logs.RetentionDays_ONE_WEEK,
//   	},
//   })
//
//   api.AddChannelNamespace(jsii.String("default"))
//
// See: https://docs.aws.amazon.com/appsync/latest/eventapi/configure-event-api-auth.html
//
type AppSyncAuthProvider struct {
	// One of possible authorization types AppSync supports.
	// Default: - `AuthorizationType.API_KEY`
	//
	AuthorizationType AppSyncAuthorizationType `field:"required" json:"authorizationType" yaml:"authorizationType"`
	// If authorizationType is `AuthorizationType.API_KEY`, this option can be configured.
	// Default: - name: 'DefaultAPIKey'.
	//
	ApiKeyConfig *AppSyncApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// If authorizationType is `AuthorizationType.USER_POOL`, this option is required.
	// Default: - none.
	//
	CognitoConfig *AppSyncCognitoConfig `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// If authorizationType is `AuthorizationType.LAMBDA`, this option is required.
	// Default: - none.
	//
	LambdaAuthorizerConfig *AppSyncLambdaAuthorizerConfig `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// If authorizationType is `AuthorizationType.OIDC`, this option is required.
	// Default: - none.
	//
	OpenIdConnectConfig *AppSyncOpenIdConnectConfig `field:"optional" json:"openIdConnectConfig" yaml:"openIdConnectConfig"`
}

