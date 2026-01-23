package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Configuration properties of an Amazon Redshift Query event.
//
// Example:
//   import redshiftserverless "github.com/aws/aws-cdk-go/awscdk"
//
//   var workgroup CfnWorkgroup
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.AddTarget(targets.NewRedshiftQuery(workgroup.attrWorkgroupWorkgroupArn, &RedshiftQueryProps{
//   	Database: jsii.String("dev"),
//   	DeadLetterQueue: dlq,
//   	Sql: []*string{
//   		jsii.String("SELECT * FROM foo"),
//   		jsii.String("SELECT * FROM baz"),
//   	},
//   }))
//
type RedshiftQueryProps struct {
	// The Amazon Redshift database to run the query against.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The SQL queries to be executed.
	//
	// Each query is run sequentially within a single transaction; the next query in the array will only execute after the previous one has successfully completed.
	//
	// - When multiple sql queries are included, this will use the `batchExecuteStatement` API. Therefore, if any statement fails, the entire transaction is rolled back.
	// - If a single SQL statement is to be executed, this will use the `executeStatement` API.
	// Default: - No SQL query is specified.
	//
	Sql *[]*string `field:"required" json:"sql" yaml:"sql"`
	// The Amazon Redshift database user to run the query as.
	//
	// This is required when authenticating via temporary credentials.
	// Default: - No Database user is specified.
	//
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The queue to be used as dead letter queue.
	// Default: - No dead letter queue is specified.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The input to the state machine execution.
	// Default: - the entire EventBridge event.
	//
	Input awsevents.RuleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The IAM role to be used to execute the SQL statement.
	// Default: - a new role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The secret containing the password for the database user.
	//
	// This is required when authenticating via Secrets Manager.
	// If the full secret ARN is not specified, this will instead use the secret name.
	// Default: - No secret is specified.
	//
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// Should an event be sent back to Event Bridge when the SQL statement is executed.
	// Default: false.
	//
	SendEventBridgeEvent *bool `field:"optional" json:"sendEventBridgeEvent" yaml:"sendEventBridgeEvent"`
	// The name of the SQL statement.
	//
	// You can name the SQL statement for identitfication purposes. If you would like Amazon Redshift to identify the Event Bridge rule, and present it in the Amazon Redshift console, append a `QS2-` prefix to the statement name.
	// Default: - No statement name is specified.
	//
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
}

