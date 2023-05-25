package awspipes


// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API BatchExecuteStatement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetRedshiftDataParametersProperty := &PipeTargetRedshiftDataParametersProperty{
//   	Database: jsii.String("database"),
//   	Sqls: []*string{
//   		jsii.String("sqls"),
//   	},
//
//   	// the properties below are optional
//   	DbUser: jsii.String("dbUser"),
//   	SecretManagerArn: jsii.String("secretManagerArn"),
//   	StatementName: jsii.String("statementName"),
//   	WithEvent: jsii.Boolean(false),
//   }
//
type CfnPipe_PipeTargetRedshiftDataParametersProperty struct {
	// The name of the database.
	//
	// Required when authenticating using temporary credentials.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The SQL statement text to run.
	Sqls *[]*string `field:"required" json:"sqls" yaml:"sqls"`
	// The database user name.
	//
	// Required when authenticating using temporary credentials.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The name or ARN of the secret that enables access to the database.
	//
	// Required when authenticating using Secrets Manager .
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the SQL statement.
	//
	// You can name the SQL statement when you create it to identify the query.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

