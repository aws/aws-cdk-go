// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha

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
//   logConfig := &logConfig{
//   	retention: logs.retentionDays_ONE_WEEK,
//   }
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	authorizationConfig: &authorizationConfig{
//   	},
//   	name: jsii.String("myApi"),
//   	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("myApi.graphql"))),
//   	logConfig: logConfig,
//   })
//
// Experimental.
type LogConfig struct {
	// exclude verbose content.
	// Experimental.
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// log level for fields.
	// Experimental.
	FieldLogLevel FieldLogLevel `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// By default AppSync keeps the logs infinitely. When updating this property,
	// unsetting it doesn't remove the log retention policy.
	// To remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	Retention awslogs.RetentionDays `field:"optional" json:"retention" yaml:"retention"`
	// The role for CloudWatch Logs.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

