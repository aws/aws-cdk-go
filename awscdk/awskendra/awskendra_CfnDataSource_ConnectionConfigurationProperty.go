package awskendra


// Provides the configuration information that's required to connect to a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionConfigurationProperty := &connectionConfigurationProperty{
//   	databaseHost: jsii.String("databaseHost"),
//   	databaseName: jsii.String("databaseName"),
//   	databasePort: jsii.Number(123),
//   	secretArn: jsii.String("secretArn"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnDataSource_ConnectionConfigurationProperty struct {
	// The name of the host for the database.
	//
	// Can be either a string (host.subdomain.domain.tld) or an IPv4 or IPv6 address.
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// The name of the database containing the document data.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The port that the database uses for connections.
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of credentials stored in AWS Secrets Manager .
	//
	// The credentials should be a user/password pair. For more information, see [Using a Database Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-database.html) . For more information about AWS Secrets Manager , see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The name of the table that contains the document data.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

