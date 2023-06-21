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
//   logConfig := &LogConfig{
//   	Retention: logs.RetentionDays_ONE_WEEK,
//   }
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	AuthorizationConfig: &AuthorizationConfig{
//   	},
//   	Name: jsii.String("myApi"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("myApi.graphql"))),
//   	LogConfig: LogConfig,
//   })
//
type LogConfig struct {
	// exclude verbose content.
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// log level for fields.
	FieldLogLevel FieldLogLevel `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// By default AppSync keeps the logs infinitely. When updating this property,
	// unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
	// The role for CloudWatch Logs.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

