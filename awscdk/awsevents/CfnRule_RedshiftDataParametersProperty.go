package awsevents


// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftDataParametersProperty := &RedshiftDataParametersProperty{
//   	Database: jsii.String("database"),
//   	Sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	DbUser: jsii.String("dbUser"),
//   	SecretManagerArn: jsii.String("secretManagerArn"),
//   	StatementName: jsii.String("statementName"),
//   	WithEvent: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html
//
type CfnRule_RedshiftDataParametersProperty struct {
	// The name of the database.
	//
	// Required when authenticating using temporary credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html#cfn-events-rule-redshiftdataparameters-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// The SQL statement text to run.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html#cfn-events-rule-redshiftdataparameters-sql
	//
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// The database user name.
	//
	// Required when authenticating using temporary credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html#cfn-events-rule-redshiftdataparameters-dbuser
	//
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The name or ARN of the secret that enables access to the database.
	//
	// Required when authenticating using AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html#cfn-events-rule-redshiftdataparameters-secretmanagerarn
	//
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the SQL statement.
	//
	// You can name the SQL statement when you create it to identify the query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html#cfn-events-rule-redshiftdataparameters-statementname
	//
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-redshiftdataparameters.html#cfn-events-rule-redshiftdataparameters-withevent
	//
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

