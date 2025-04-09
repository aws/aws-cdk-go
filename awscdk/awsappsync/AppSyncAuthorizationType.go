package awsappsync


// enum with all possible values for AppSync authorization type.
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
type AppSyncAuthorizationType string

const (
	// API Key authorization type.
	AppSyncAuthorizationType_API_KEY AppSyncAuthorizationType = "API_KEY"
	// AWS IAM authorization type.
	//
	// Can be used with Cognito Identity Pool federated credentials.
	AppSyncAuthorizationType_IAM AppSyncAuthorizationType = "IAM"
	// Cognito User Pool authorization type.
	AppSyncAuthorizationType_USER_POOL AppSyncAuthorizationType = "USER_POOL"
	// OpenID Connect authorization type.
	AppSyncAuthorizationType_OIDC AppSyncAuthorizationType = "OIDC"
	// Lambda authorization type.
	AppSyncAuthorizationType_LAMBDA AppSyncAuthorizationType = "LAMBDA"
)

