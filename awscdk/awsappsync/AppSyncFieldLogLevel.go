package awsappsync


// log-level for fields in AppSync.
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
type AppSyncFieldLogLevel string

const (
	// Resolver logging is disabled.
	AppSyncFieldLogLevel_NONE AppSyncFieldLogLevel = "NONE"
	// Only Error messages appear in logs.
	AppSyncFieldLogLevel_ERROR AppSyncFieldLogLevel = "ERROR"
	// Info and Error messages appear in logs.
	AppSyncFieldLogLevel_INFO AppSyncFieldLogLevel = "INFO"
	// Debug, Info, and Error messages, appear in logs.
	AppSyncFieldLogLevel_DEBUG AppSyncFieldLogLevel = "DEBUG"
	// All messages (Debug, Error, Info, and Trace) appear in logs.
	AppSyncFieldLogLevel_ALL AppSyncFieldLogLevel = "ALL"
)

