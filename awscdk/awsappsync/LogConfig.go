package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// Logging configuration for AppSync.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	AuthorizationConfig: &AuthorizationConfig{
//   	},
//   	Name: jsii.String("myApi"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("myApi.graphql"))),
//   	LogConfig: &LogConfig{
//   		FieldLogLevel: appsync.FieldLogLevel_INFO,
//   		Retention: logs.RetentionDays_ONE_WEEK,
//   	},
//   })
//
type LogConfig struct {
	// exclude verbose content.
	// Default: false.
	//
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// log level for fields.
	// Default: - Use AppSync default.
	//
	FieldLogLevel FieldLogLevel `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
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
	Role interfacesawsiam.IRoleRef `field:"optional" json:"role" yaml:"role"`
}

