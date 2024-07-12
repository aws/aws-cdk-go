package awsdms


// Provides information that defines an Oracle endpoint.
//
// This information includes the output format of records applied to the endpoint and details of transaction and control table data information. For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oracleSettingsProperty := &OracleSettingsProperty{
//   	AccessAlternateDirectly: jsii.Boolean(false),
//   	AdditionalArchivedLogDestId: jsii.Number(123),
//   	AddSupplementalLogging: jsii.Boolean(false),
//   	AllowSelectNestedTables: jsii.Boolean(false),
//   	ArchivedLogDestId: jsii.Number(123),
//   	ArchivedLogsOnly: jsii.Boolean(false),
//   	AsmPassword: jsii.String("asmPassword"),
//   	AsmServer: jsii.String("asmServer"),
//   	AsmUser: jsii.String("asmUser"),
//   	CharLengthSemantics: jsii.String("charLengthSemantics"),
//   	DirectPathNoLog: jsii.Boolean(false),
//   	DirectPathParallelLoad: jsii.Boolean(false),
//   	EnableHomogenousTablespace: jsii.Boolean(false),
//   	ExtraArchivedLogDestIds: []interface{}{
//   		jsii.Number(123),
//   	},
//   	FailTasksOnLobTruncation: jsii.Boolean(false),
//   	NumberDatatypeScale: jsii.Number(123),
//   	OraclePathPrefix: jsii.String("oraclePathPrefix"),
//   	ParallelAsmReadThreads: jsii.Number(123),
//   	ReadAheadBlocks: jsii.Number(123),
//   	ReadTableSpaceName: jsii.Boolean(false),
//   	ReplacePathPrefix: jsii.Boolean(false),
//   	RetryInterval: jsii.Number(123),
//   	SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   	SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   	SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   	SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	SecurityDbEncryption: jsii.String("securityDbEncryption"),
//   	SecurityDbEncryptionName: jsii.String("securityDbEncryptionName"),
//   	SpatialDataOptionToGeoJsonFunctionName: jsii.String("spatialDataOptionToGeoJsonFunctionName"),
//   	StandbyDelayTime: jsii.Number(123),
//   	UseAlternateFolderForOnline: jsii.Boolean(false),
//   	UseBFile: jsii.Boolean(false),
//   	UseDirectPathFullLoad: jsii.Boolean(false),
//   	UseLogminerReader: jsii.Boolean(false),
//   	UsePathPrefix: jsii.String("usePathPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html
//
type CfnEndpoint_OracleSettingsProperty struct {
	// Set this attribute to `false` in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This tells the DMS instance to not access redo logs through any specified path prefix replacement using direct file access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-accessalternatedirectly
	//
	AccessAlternateDirectly interface{} `field:"optional" json:"accessAlternateDirectly" yaml:"accessAlternateDirectly"`
	// Set this attribute with `ArchivedLogDestId` in a primary/ standby setup.
	//
	// This attribute is useful in the case of a switchover. In this case, AWS DMS needs to know which destination to get archive redo logs from to read changes. This need arises because the previous primary instance is now a standby instance after switchover.
	//
	// Although AWS DMS supports the use of the Oracle `RESETLOGS` option to open the database, never use `RESETLOGS` unless necessary. For additional information about `RESETLOGS` , see [RMAN Data Repair Concepts](https://docs.aws.amazon.com/https://docs.oracle.com/en/database/oracle/oracle-database/19/bradv/rman-data-repair-concepts.html#GUID-1805CCF7-4AF2-482D-B65A-998192F89C2B) in the *Oracle Database Backup and Recovery User's Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-additionalarchivedlogdestid
	//
	AdditionalArchivedLogDestId *float64 `field:"optional" json:"additionalArchivedLogDestId" yaml:"additionalArchivedLogDestId"`
	// Set this attribute to set up table-level supplemental logging for the Oracle database.
	//
	// This attribute enables PRIMARY KEY supplemental logging on all tables selected for a migration task.
	//
	// If you use this option, you still need to enable database-level supplemental logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-addsupplementallogging
	//
	AddSupplementalLogging interface{} `field:"optional" json:"addSupplementalLogging" yaml:"addSupplementalLogging"`
	// Set this attribute to `true` to enable replication of Oracle tables containing columns that are nested tables or defined types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-allowselectnestedtables
	//
	AllowSelectNestedTables interface{} `field:"optional" json:"allowSelectNestedTables" yaml:"allowSelectNestedTables"`
	// Specifies the ID of the destination for the archived redo logs.
	//
	// This value should be the same as a number in the dest_id column of the v$archived_log view. If you work with an additional redo log destination, use the `AdditionalArchivedLogDestId` option to specify the additional destination ID. Doing this improves performance by ensuring that the correct logs are accessed from the outset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-archivedlogdestid
	//
	ArchivedLogDestId *float64 `field:"optional" json:"archivedLogDestId" yaml:"archivedLogDestId"`
	// When this field is set to `True` , AWS DMS only accesses the archived redo logs.
	//
	// If the archived redo logs are stored on Automatic Storage Management (ASM) only, the AWS DMS user account needs to be granted ASM privileges.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-archivedlogsonly
	//
	ArchivedLogsOnly interface{} `field:"optional" json:"archivedLogsOnly" yaml:"archivedLogsOnly"`
	// For an Oracle source endpoint, your Oracle Automatic Storage Management (ASM) password.
	//
	// You can set this value from the `*asm_user_password*` value. You set this value as part of the comma-separated value that you set to the `Password` request parameter when you create the endpoint to access transaction logs using Binary Reader. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmpassword
	//
	AsmPassword *string `field:"optional" json:"asmPassword" yaml:"asmPassword"`
	// For an Oracle source endpoint, your ASM server address.
	//
	// You can set this value from the `asm_server` value. You set `asm_server` as part of the extra connection attribute string to access an Oracle server with Binary Reader that uses ASM. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmserver
	//
	AsmServer *string `field:"optional" json:"asmServer" yaml:"asmServer"`
	// For an Oracle source endpoint, your ASM user name.
	//
	// You can set this value from the `asm_user` value. You set `asm_user` as part of the extra connection attribute string to access an Oracle server with Binary Reader that uses ASM. For more information, see [Configuration for change data capture (CDC) on an Oracle source database](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC.Configuration) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-asmuser
	//
	AsmUser *string `field:"optional" json:"asmUser" yaml:"asmUser"`
	// Specifies whether the length of a character column is in bytes or in characters.
	//
	// To indicate that the character column length is in characters, set this attribute to `CHAR` . Otherwise, the character column length is in bytes.
	//
	// Example: `charLengthSemantics=CHAR;`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-charlengthsemantics
	//
	CharLengthSemantics *string `field:"optional" json:"charLengthSemantics" yaml:"charLengthSemantics"`
	// When set to `true` , this attribute helps to increase the commit rate on the Oracle target database by writing directly to tables and not writing a trail to database logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-directpathnolog
	//
	DirectPathNoLog interface{} `field:"optional" json:"directPathNoLog" yaml:"directPathNoLog"`
	// When set to `true` , this attribute specifies a parallel load when `useDirectPathFullLoad` is set to `Y` .
	//
	// This attribute also only applies when you use the AWS DMS parallel load feature. Note that the target table cannot have any constraints or indexes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-directpathparallelload
	//
	DirectPathParallelLoad interface{} `field:"optional" json:"directPathParallelLoad" yaml:"directPathParallelLoad"`
	// Set this attribute to enable homogenous tablespace replication and create existing tables or indexes under the same tablespace on the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-enablehomogenoustablespace
	//
	EnableHomogenousTablespace interface{} `field:"optional" json:"enableHomogenousTablespace" yaml:"enableHomogenousTablespace"`
	// Specifies the IDs of one more destinations for one or more archived redo logs.
	//
	// These IDs are the values of the `dest_id` column in the `v$archived_log` view. Use this setting with the `archivedLogDestId` extra connection attribute in a primary-to-single setup or a primary-to-multiple-standby setup.
	//
	// This setting is useful in a switchover when you use an Oracle Data Guard database as a source. In this case, AWS DMS needs information about what destination to get archive redo logs from to read changes. AWS DMS needs this because after the switchover the previous primary is a standby instance. For example, in a primary-to-single standby setup you might apply the following settings.
	//
	// `archivedLogDestId=1; ExtraArchivedLogDestIds=[2]`
	//
	// In a primary-to-multiple-standby setup, you might apply the following settings.
	//
	// `archivedLogDestId=1; ExtraArchivedLogDestIds=[2,3,4]`
	//
	// Although AWS DMS supports the use of the Oracle `RESETLOGS` option to open the database, never use `RESETLOGS` unless it's necessary. For more information about `RESETLOGS` , see [RMAN Data Repair Concepts](https://docs.aws.amazon.com/https://docs.oracle.com/en/database/oracle/oracle-database/19/bradv/rman-data-repair-concepts.html#GUID-1805CCF7-4AF2-482D-B65A-998192F89C2B) in the *Oracle Database Backup and Recovery User's Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-extraarchivedlogdestids
	//
	ExtraArchivedLogDestIds interface{} `field:"optional" json:"extraArchivedLogDestIds" yaml:"extraArchivedLogDestIds"`
	// When set to `true` , this attribute causes a task to fail if the actual size of an LOB column is greater than the specified `LobMaxSize` .
	//
	// If a task is set to limited LOB mode and this option is set to `true` , the task fails instead of truncating the LOB data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-failtasksonlobtruncation
	//
	FailTasksOnLobTruncation interface{} `field:"optional" json:"failTasksOnLobTruncation" yaml:"failTasksOnLobTruncation"`
	// Specifies the number scale.
	//
	// You can select a scale up to 38, or you can select FLOAT. By default, the NUMBER data type is converted to precision 38, scale 10.
	//
	// Example: `numberDataTypeScale=12`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-numberdatatypescale
	//
	NumberDatatypeScale *float64 `field:"optional" json:"numberDatatypeScale" yaml:"numberDatatypeScale"`
	// Set this string attribute to the required value in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This value specifies the default Oracle root used to access the redo logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-oraclepathprefix
	//
	OraclePathPrefix *string `field:"optional" json:"oraclePathPrefix" yaml:"oraclePathPrefix"`
	// Set this attribute to change the number of threads that DMS configures to perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).
	//
	// You can specify an integer value between 2 (the default) and 8 (the maximum). Use this attribute together with the `readAheadBlocks` attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-parallelasmreadthreads
	//
	ParallelAsmReadThreads *float64 `field:"optional" json:"parallelAsmReadThreads" yaml:"parallelAsmReadThreads"`
	// Set this attribute to change the number of read-ahead blocks that DMS configures to perform a change data capture (CDC) load using Oracle Automatic Storage Management (ASM).
	//
	// You can specify an integer value between 1000 (the default) and 200,000 (the maximum).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-readaheadblocks
	//
	ReadAheadBlocks *float64 `field:"optional" json:"readAheadBlocks" yaml:"readAheadBlocks"`
	// When set to `true` , this attribute supports tablespace replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-readtablespacename
	//
	ReadTableSpaceName interface{} `field:"optional" json:"readTableSpaceName" yaml:"readTableSpaceName"`
	// Set this attribute to true in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This setting tells DMS instance to replace the default Oracle root with the specified `usePathPrefix` setting to access the redo logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-replacepathprefix
	//
	ReplacePathPrefix interface{} `field:"optional" json:"replacePathPrefix" yaml:"replacePathPrefix"`
	// Specifies the number of seconds that the system waits before resending a query.
	//
	// Example: `retryInterval=6;`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-retryinterval
	//
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// The full Amazon Resource Name (ARN) of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the value in `SecretsManagerSecret` .
	//
	// The role must allow the `iam:PassRole` action. `SecretsManagerSecret` has the value of the AWS Secrets Manager secret that allows access to the Oracle endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerSecretId` . Or you can specify clear-text values for `UserName` , `Password` , `ServerName` , and `Port` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerSecret` , the corresponding `SecretsManagerAccessRoleArn` , and the `SecretsManagerSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageraccessrolearn
	//
	SecretsManagerAccessRoleArn *string `field:"optional" json:"secretsManagerAccessRoleArn" yaml:"secretsManagerAccessRoleArn"`
	// Required only if your Oracle endpoint uses Advanced Storage Manager (ASM).
	//
	// The full ARN of the IAM role that specifies AWS DMS as the trusted entity and grants the required permissions to access the `SecretsManagerOracleAsmSecret` . This `SecretsManagerOracleAsmSecret` has the secret value that allows access to the Oracle ASM of the endpoint.
	//
	// > You can specify one of two sets of values for these permissions. You can specify the values for this setting and `SecretsManagerOracleAsmSecretId` . Or you can specify clear-text values for `AsmUser` , `AsmPassword` , and `AsmServerName` . You can't specify both.
	// >
	// > For more information on creating this `SecretsManagerOracleAsmSecret` , the corresponding `SecretsManagerOracleAsmAccessRoleArn` , and the `SecretsManagerOracleAsmSecretId` that is required to access it, see [Using secrets to access AWS Database Migration Service resources](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.html#security-iam-secretsmanager) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageroracleasmaccessrolearn
	//
	SecretsManagerOracleAsmAccessRoleArn *string `field:"optional" json:"secretsManagerOracleAsmAccessRoleArn" yaml:"secretsManagerOracleAsmAccessRoleArn"`
	// Required only if your Oracle endpoint uses Advanced Storage Manager (ASM).
	//
	// The full ARN, partial ARN, or display name of the `SecretsManagerOracleAsmSecret` that contains the Oracle ASM connection details for the Oracle endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanageroracleasmsecretid
	//
	SecretsManagerOracleAsmSecretId *string `field:"optional" json:"secretsManagerOracleAsmSecretId" yaml:"secretsManagerOracleAsmSecretId"`
	// The full ARN, partial ARN, or display name of the `SecretsManagerSecret` that contains the Oracle endpoint connection details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-secretsmanagersecretid
	//
	SecretsManagerSecretId *string `field:"optional" json:"secretsManagerSecretId" yaml:"secretsManagerSecretId"`
	// For an Oracle source endpoint, the transparent data encryption (TDE) password required by AWM DMS to access Oracle redo logs encrypted by TDE using Binary Reader.
	//
	// It is also the `*TDE_Password*` part of the comma-separated value you set to the `Password` request parameter when you create the endpoint. The `SecurityDbEncryptian` setting is related to this `SecurityDbEncryptionName` setting. For more information, see [Supported encryption methods for using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.Encryption) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-securitydbencryption
	//
	SecurityDbEncryption *string `field:"optional" json:"securityDbEncryption" yaml:"securityDbEncryption"`
	// For an Oracle source endpoint, the name of a key used for the transparent data encryption (TDE) of the columns and tablespaces in an Oracle source database that is encrypted using TDE.
	//
	// The key value is the value of the `SecurityDbEncryption` setting. For more information on setting the key name value of `SecurityDbEncryptionName` , see the information and example for setting the `securityDbEncryptionName` extra connection attribute in [Supported encryption methods for using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.Encryption) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-securitydbencryptionname
	//
	SecurityDbEncryptionName *string `field:"optional" json:"securityDbEncryptionName" yaml:"securityDbEncryptionName"`
	// Use this attribute to convert `SDO_GEOMETRY` to `GEOJSON` format.
	//
	// By default, DMS calls the `SDO2GEOJSON` custom function if present and accessible. Or you can create your own custom function that mimics the operation of `SDOGEOJSON` and set `SpatialDataOptionToGeoJsonFunctionName` to call it instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-spatialdataoptiontogeojsonfunctionname
	//
	SpatialDataOptionToGeoJsonFunctionName *string `field:"optional" json:"spatialDataOptionToGeoJsonFunctionName" yaml:"spatialDataOptionToGeoJsonFunctionName"`
	// Use this attribute to specify a time in minutes for the delay in standby sync.
	//
	// If the source is an Oracle Active Data Guard standby database, use this attribute to specify the time lag between primary and standby databases.
	//
	// In AWS DMS , you can create an Oracle CDC task that uses an Active Data Guard standby instance as a source for replicating ongoing changes. Doing this eliminates the need to connect to an active database that might be in production.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-standbydelaytime
	//
	StandbyDelayTime *float64 `field:"optional" json:"standbyDelayTime" yaml:"standbyDelayTime"`
	// Set this attribute to `true` in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This tells the DMS instance to use any specified prefix replacement to access all online redo logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usealternatefolderforonline
	//
	UseAlternateFolderForOnline interface{} `field:"optional" json:"useAlternateFolderForOnline" yaml:"useAlternateFolderForOnline"`
	// Set this attribute to True to capture change data using the Binary Reader utility.
	//
	// Set `UseLogminerReader` to False to set this attribute to True. To use Binary Reader with Amazon RDS for Oracle as the source, you set additional attributes. For more information about using this setting with Oracle Automatic Storage Management (ASM), see [Using Oracle LogMiner or AWS DMS Binary Reader for CDC](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usebfile
	//
	UseBFile interface{} `field:"optional" json:"useBFile" yaml:"useBFile"`
	// Set this attribute to True to have AWS DMS use a direct path full load.
	//
	// Specify this value to use the direct path protocol in the Oracle Call Interface (OCI). By using this OCI protocol, you can bulk-load Oracle target tables during a full load.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usedirectpathfullload
	//
	UseDirectPathFullLoad interface{} `field:"optional" json:"useDirectPathFullLoad" yaml:"useDirectPathFullLoad"`
	// Set this attribute to True to capture change data using the Oracle LogMiner utility (the default).
	//
	// Set this attribute to False if you want to access the redo logs as a binary file. When you set `UseLogminerReader` to False, also set `UseBfile` to True. For more information on this setting and using Oracle ASM, see [Using Oracle LogMiner or AWS DMS Binary Reader for CDC](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.CDC) in the *AWS DMS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-uselogminerreader
	//
	UseLogminerReader interface{} `field:"optional" json:"useLogminerReader" yaml:"useLogminerReader"`
	// Set this string attribute to the required value in order to use the Binary Reader to capture change data for an Amazon RDS for Oracle as the source.
	//
	// This value specifies the path prefix used to replace the default Oracle root to access the redo logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-oraclesettings.html#cfn-dms-endpoint-oraclesettings-usepathprefix
	//
	UsePathPrefix *string `field:"optional" json:"usePathPrefix" yaml:"usePathPrefix"`
}

