package awsappsync


// Authorization configuration for the Event API.
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
type EventApiAuthConfig struct {
	// Auth providers for use in connection, publish, and subscribe operations.
	// Default: - API Key authorization.
	//
	AuthProviders *[]*AppSyncAuthProvider `field:"optional" json:"authProviders" yaml:"authProviders"`
	// Connection auth modes.
	// Default: - API Key authorization.
	//
	ConnectionAuthModeTypes *[]AppSyncAuthorizationType `field:"optional" json:"connectionAuthModeTypes" yaml:"connectionAuthModeTypes"`
	// Default publish auth modes.
	// Default: - API Key authorization.
	//
	DefaultPublishAuthModeTypes *[]AppSyncAuthorizationType `field:"optional" json:"defaultPublishAuthModeTypes" yaml:"defaultPublishAuthModeTypes"`
	// Default subscribe auth modes.
	// Default: - API Key authorization.
	//
	DefaultSubscribeAuthModeTypes *[]AppSyncAuthorizationType `field:"optional" json:"defaultSubscribeAuthModeTypes" yaml:"defaultSubscribeAuthModeTypes"`
}

