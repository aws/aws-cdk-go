package awsevents


// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftDataParametersProperty := &redshiftDataParametersProperty{
//   	database: jsii.String("database"),
//   	sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	dbUser: jsii.String("dbUser"),
//   	secretManagerArn: jsii.String("secretManagerArn"),
//   	statementName: jsii.String("statementName"),
//   	withEvent: jsii.Boolean(false),
//   }
//
type CfnRule_RedshiftDataParametersProperty struct {
	// The name of the database.
	//
	// Required when authenticating using temporary credentials.
	Database *string `field:"required" json:"database" yaml:"database"`
	// The SQL statement text to run.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// The database user name.
	//
	// Required when authenticating using temporary credentials.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The name or ARN of the secret that enables access to the database.
	//
	// Required when authenticating using AWS Secrets Manager.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the SQL statement.
	//
	// You can name the SQL statement when you create it to identify the query.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

