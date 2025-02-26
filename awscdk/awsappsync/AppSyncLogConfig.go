package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Logging configuration for AppSync.
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
type AppSyncLogConfig struct {
	// exclude verbose content.
	// Default: false.
	//
	ExcludeVerboseContent *bool `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// log level for fields.
	// Default: - Use AppSync default.
	//
	FieldLogLevel AppSyncFieldLogLevel `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// By default AppSync keeps the logs infinitely. When updating this property,
	// unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	// Default: RetentionDays.INFINITE
	//
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
	// The role for CloudWatch Logs.
	// Default: - None.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

