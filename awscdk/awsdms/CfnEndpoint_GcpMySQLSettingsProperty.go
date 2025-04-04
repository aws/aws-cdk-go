package awsdms


// Provides information that defines a GCP MySQL endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gcpMySQLSettingsProperty := &GcpMySQLSettingsProperty{
//   	AfterConnectScript: jsii.String("afterConnectScript"),
//   	CleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   	DatabaseName: jsii.String("databaseName"),
//   	EventsPollInterval: jsii.Number(123),
//   	MaxFileSize: jsii.Number(123),
//   	ParallelLoadThreads: jsii.Number(123),
//   	Password: jsii.String("password"),
//   	Port: jsii.Number(123),
//   	SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	ServerName: jsii.String("serverName"),
//   	ServerTimezone: jsii.String("serverTimezone"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html
//
type CfnEndpoint_GcpMySQLSettingsProperty struct {
	// Specifies a script to run immediately after AWS DMS connects to the endpoint.
	//
	// The migration task continues running regardless if the SQL statement succeeds or fails.
	//
	// For this parameter, provide the code of the script itself, not the name of a file containing the script.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-afterconnectscript
	//
	AfterConnectScript *string `field:"optional" json:"afterConnectScript" yaml:"afterConnectScript"`
	// Adjusts the behavior of AWS DMS when migrating from an SQL Server source database that is hosted as part of an Always On availability group cluster.
	//
	// If you need AWS DMS to poll all the nodes in the Always On cluster for transaction backups, set this attribute to `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-cleansourcemetadataonmismatch
	//
	CleanSourceMetadataOnMismatch interface{} `field:"optional" json:"cleanSourceMetadataOnMismatch" yaml:"cleanSourceMetadataOnMismatch"`
	// Database name for the endpoint.
	//
	// For a MySQL source or target endpoint, don't explicitly specify the database using the `DatabaseName` request parameter on either the `CreateEndpoint` or `ModifyEndpoint` API call. Specifying `DatabaseName` when you create or modify a MySQL endpoint replicates all the task tables to this single database. For MySQL endpoints, you specify the database only when you specify the schema in the table-mapping rules of the AWS DMS task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Specifies how often to check the binary log for new changes/events when the database is idle.
	//
	// The default is five seconds.
	//
	// Example: `eventsPollInterval=5;`
	//
	// In the example, AWS DMS checks for changes in the binary logs every five seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-eventspollinterval
	//
	EventsPollInterval *float64 `field:"optional" json:"eventsPollInterval" yaml:"eventsPollInterval"`
	// Specifies the maximum size (in KB) of any .csv file used to transfer data to a MySQL-compatible database.
	//
	// Example: `maxFileSize=512`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-maxfilesize
	//
	MaxFileSize *float64 `field:"optional" json:"maxFileSize" yaml:"maxFileSize"`
	// Improves performance when loading data into the MySQL-compatible target database.
	//
	// Specifies how many threads to use to load the data into the MySQL-compatible target database. Setting a large number of threads can have an adverse effect on database performance, because a separate connection is required for each thread. The default is one.
	//
	// Example: `parallelLoadThreads=1`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-parallelloadthreads
	//
	ParallelLoadThreads *float64 `field:"optional" json:"parallelLoadThreads" yaml:"parallelLoadThreads"`
	// Endpoint connection password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret.` The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the MySQL endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-secretsmanageraccessrolearn
	//
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MySQL endpoint connection details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-secretsmanagersecretid
	//
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// The MySQL host name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Specifies the time zone for the source MySQL database. Don't enclose time zones in single quotation marks.
	//
	// Example: `serverTimezone=US/Pacific;`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-servertimezone
	//
	ServerTimezone *string `field:"optional" json:"serverTimezone" yaml:"serverTimezone"`
	// Endpoint connection user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-gcpmysqlsettings.html#cfn-dms-endpoint-gcpmysqlsettings-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

