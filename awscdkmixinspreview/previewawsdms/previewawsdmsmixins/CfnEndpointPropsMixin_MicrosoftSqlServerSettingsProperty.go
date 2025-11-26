package previewawsdmsmixins


// Provides information that defines a Microsoft SQL Server endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   microsoftSqlServerSettingsProperty := &MicrosoftSqlServerSettingsProperty{
//   	BcpPacketSize: jsii.Number(123),
//   	ControlTablesFileGroup: jsii.String("controlTablesFileGroup"),
//   	DatabaseName: jsii.String("databaseName"),
//   	ForceLobLookup: jsii.Boolean(false),
//   	Password: jsii.String("password"),
//   	Port: jsii.Number(123),
//   	QuerySingleAlwaysOnNode: jsii.Boolean(false),
//   	ReadBackupOnly: jsii.Boolean(false),
//   	SafeguardPolicy: jsii.String("safeguardPolicy"),
//   	SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	ServerName: jsii.String("serverName"),
//   	TlogAccessMode: jsii.String("tlogAccessMode"),
//   	TrimSpaceInChar: jsii.Boolean(false),
//   	UseBcpFullLoad: jsii.Boolean(false),
//   	Username: jsii.String("username"),
//   	UseThirdPartyBackupDevice: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html
//
type CfnEndpointPropsMixin_MicrosoftSqlServerSettingsProperty struct {
	// The maximum size of the packets (in bytes) used to transfer data using BCP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-bcppacketsize
	//
	BcpPacketSize *float64 `field:"optional" json:"bcpPacketSize" yaml:"bcpPacketSize"`
	// Specifies a file group for the AWS DMS internal tables.
	//
	// When the replication task starts, all the internal AWS DMS control tables (awsdms_ apply_exception, awsdms_apply, awsdms_changes) are created for the specified file group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-controltablesfilegroup
	//
	ControlTablesFileGroup *string `field:"optional" json:"controlTablesFileGroup" yaml:"controlTablesFileGroup"`
	// Database name for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Forces LOB lookup on inline LOB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-forceloblookup
	//
	ForceLobLookup interface{} `field:"optional" json:"forceLobLookup" yaml:"forceLobLookup"`
	// Endpoint connection password.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Endpoint TCP port.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Cleans and recreates table metadata information on the replication instance when a mismatch occurs.
	//
	// An example is a situation where running an alter DDL statement on a table might result in different information about the table cached in the replication instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-querysinglealwaysonnode
	//
	QuerySingleAlwaysOnNode interface{} `field:"optional" json:"querySingleAlwaysOnNode" yaml:"querySingleAlwaysOnNode"`
	// When this attribute is set to `Y` , AWS DMS only reads changes from transaction log backups and doesn't read from the active transaction log file during ongoing replication.
	//
	// Setting this parameter to `Y` enables you to control active transaction log file growth during full load and ongoing replication tasks. However, it can add some source latency to ongoing replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-readbackuponly
	//
	ReadBackupOnly interface{} `field:"optional" json:"readBackupOnly" yaml:"readBackupOnly"`
	// Use this attribute to minimize the need to access the backup log and enable AWS DMS to prevent truncation using one of the following two methods.
	//
	// *Start transactions in the database:* This is the default method. When this method is used, AWS DMS prevents TLOG truncation by mimicking a transaction in the database. As long as such a transaction is open, changes that appear after the transaction started aren't truncated. If you need Microsoft Replication to be enabled in your database, then you must choose this method.
	//
	// *Exclusively use sp_repldone within a single task* : When this method is used, AWS DMS reads the changes and then uses sp_repldone to mark the TLOG transactions as ready for truncation. Although this method doesn't involve any transactional activities, it can only be used when Microsoft Replication isn't running. Also, when using this method, only one AWS DMS task can access the database at any given time. Therefore, if you need to run parallel AWS DMS tasks against the same database, use the default method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-safeguardpolicy
	//
	SafeguardPolicy *string `field:"optional" json:"safeguardPolicy" yaml:"safeguardPolicy"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the SQL Server endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-secretsmanageraccessrolearn
	//
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the MicrosoftSQLServer endpoint connection details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-secretsmanagersecretid
	//
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// Fully qualified domain name of the endpoint.
	//
	// For an Amazon RDS SQL Server instance, this is the output of [DescribeDBInstances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBInstances.html) , in the `[Endpoint](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Endpoint.html) .Address` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Indicates the mode used to fetch CDC data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-tlogaccessmode
	//
	TlogAccessMode *string `field:"optional" json:"tlogAccessMode" yaml:"tlogAccessMode"`
	// Use the `TrimSpaceInChar` source endpoint setting to right-trim data on CHAR and NCHAR data types during migration.
	//
	// Setting `TrimSpaceInChar` does not left-trim data. The default value is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-trimspaceinchar
	//
	TrimSpaceInChar interface{} `field:"optional" json:"trimSpaceInChar" yaml:"trimSpaceInChar"`
	// Use this to attribute to transfer data for full-load operations using BCP.
	//
	// When the target table contains an identity column that does not exist in the source table, you must disable the use BCP for loading table option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-usebcpfullload
	//
	UseBcpFullLoad interface{} `field:"optional" json:"useBcpFullLoad" yaml:"useBcpFullLoad"`
	// Endpoint connection user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
	// When this attribute is set to `Y` , DMS processes third-party transaction log backups if they are created in native format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-microsoftsqlserversettings.html#cfn-dms-endpoint-microsoftsqlserversettings-usethirdpartybackupdevice
	//
	UseThirdPartyBackupDevice interface{} `field:"optional" json:"useThirdPartyBackupDevice" yaml:"useThirdPartyBackupDevice"`
}

