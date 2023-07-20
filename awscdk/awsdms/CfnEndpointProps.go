package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &CfnEndpointProps{
//   	EndpointType: jsii.String("endpointType"),
//   	EngineName: jsii.String("engineName"),
//
//   	// the properties below are optional
//   	CertificateArn: jsii.String("certificateArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DocDbSettings: &DocDbSettingsProperty{
//   		DocsToInvestigate: jsii.Number(123),
//   		ExtractDocId: jsii.Boolean(false),
//   		NestingLevel: jsii.String("nestingLevel"),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	DynamoDbSettings: &DynamoDbSettingsProperty{
//   		ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	ElasticsearchSettings: &ElasticsearchSettingsProperty{
//   		EndpointUri: jsii.String("endpointUri"),
//   		ErrorRetryDuration: jsii.Number(123),
//   		FullLoadErrorPercentage: jsii.Number(123),
//   		ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	EndpointIdentifier: jsii.String("endpointIdentifier"),
//   	ExtraConnectionAttributes: jsii.String("extraConnectionAttributes"),
//   	GcpMySqlSettings: &GcpMySQLSettingsProperty{
//   		AfterConnectScript: jsii.String("afterConnectScript"),
//   		CleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		DatabaseName: jsii.String("databaseName"),
//   		EventsPollInterval: jsii.Number(123),
//   		MaxFileSize: jsii.Number(123),
//   		ParallelLoadThreads: jsii.Number(123),
//   		Password: jsii.String("password"),
//   		Port: jsii.Number(123),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		ServerName: jsii.String("serverName"),
//   		ServerTimezone: jsii.String("serverTimezone"),
//   		Username: jsii.String("username"),
//   	},
//   	IbmDb2Settings: &IbmDb2SettingsProperty{
//   		CurrentLsn: jsii.String("currentLsn"),
//   		MaxKBytesPerRead: jsii.Number(123),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		SetDataCaptureChanges: jsii.Boolean(false),
//   	},
//   	KafkaSettings: &KafkaSettingsProperty{
//   		Broker: jsii.String("broker"),
//   		IncludeControlDetails: jsii.Boolean(false),
//   		IncludeNullAndEmpty: jsii.Boolean(false),
//   		IncludePartitionValue: jsii.Boolean(false),
//   		IncludeTableAlterOperations: jsii.Boolean(false),
//   		IncludeTransactionDetails: jsii.Boolean(false),
//   		MessageFormat: jsii.String("messageFormat"),
//   		MessageMaxBytes: jsii.Number(123),
//   		NoHexPrefix: jsii.Boolean(false),
//   		PartitionIncludeSchemaTable: jsii.Boolean(false),
//   		SaslPassword: jsii.String("saslPassword"),
//   		SaslUserName: jsii.String("saslUserName"),
//   		SecurityProtocol: jsii.String("securityProtocol"),
//   		SslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		SslClientCertificateArn: jsii.String("sslClientCertificateArn"),
//   		SslClientKeyArn: jsii.String("sslClientKeyArn"),
//   		SslClientKeyPassword: jsii.String("sslClientKeyPassword"),
//   		Topic: jsii.String("topic"),
//   	},
//   	KinesisSettings: &KinesisSettingsProperty{
//   		IncludeControlDetails: jsii.Boolean(false),
//   		IncludeNullAndEmpty: jsii.Boolean(false),
//   		IncludePartitionValue: jsii.Boolean(false),
//   		IncludeTableAlterOperations: jsii.Boolean(false),
//   		IncludeTransactionDetails: jsii.Boolean(false),
//   		MessageFormat: jsii.String("messageFormat"),
//   		NoHexPrefix: jsii.Boolean(false),
//   		PartitionIncludeSchemaTable: jsii.Boolean(false),
//   		ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MicrosoftSqlServerSettings: &MicrosoftSqlServerSettingsProperty{
//   		BcpPacketSize: jsii.Number(123),
//   		ControlTablesFileGroup: jsii.String("controlTablesFileGroup"),
//   		QuerySingleAlwaysOnNode: jsii.Boolean(false),
//   		ReadBackupOnly: jsii.Boolean(false),
//   		SafeguardPolicy: jsii.String("safeguardPolicy"),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		UseBcpFullLoad: jsii.Boolean(false),
//   		UseThirdPartyBackupDevice: jsii.Boolean(false),
//   	},
//   	MongoDbSettings: &MongoDbSettingsProperty{
//   		AuthMechanism: jsii.String("authMechanism"),
//   		AuthSource: jsii.String("authSource"),
//   		AuthType: jsii.String("authType"),
//   		DatabaseName: jsii.String("databaseName"),
//   		DocsToInvestigate: jsii.String("docsToInvestigate"),
//   		ExtractDocId: jsii.String("extractDocId"),
//   		NestingLevel: jsii.String("nestingLevel"),
//   		Password: jsii.String("password"),
//   		Port: jsii.Number(123),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		ServerName: jsii.String("serverName"),
//   		Username: jsii.String("username"),
//   	},
//   	MySqlSettings: &MySqlSettingsProperty{
//   		AfterConnectScript: jsii.String("afterConnectScript"),
//   		CleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		EventsPollInterval: jsii.Number(123),
//   		MaxFileSize: jsii.Number(123),
//   		ParallelLoadThreads: jsii.Number(123),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		ServerTimezone: jsii.String("serverTimezone"),
//   		TargetDbType: jsii.String("targetDbType"),
//   	},
//   	NeptuneSettings: &NeptuneSettingsProperty{
//   		ErrorRetryDuration: jsii.Number(123),
//   		IamAuthEnabled: jsii.Boolean(false),
//   		MaxFileSize: jsii.Number(123),
//   		MaxRetryCount: jsii.Number(123),
//   		S3BucketFolder: jsii.String("s3BucketFolder"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   		ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	OracleSettings: &OracleSettingsProperty{
//   		AccessAlternateDirectly: jsii.Boolean(false),
//   		AdditionalArchivedLogDestId: jsii.Number(123),
//   		AddSupplementalLogging: jsii.Boolean(false),
//   		AllowSelectNestedTables: jsii.Boolean(false),
//   		ArchivedLogDestId: jsii.Number(123),
//   		ArchivedLogsOnly: jsii.Boolean(false),
//   		AsmPassword: jsii.String("asmPassword"),
//   		AsmServer: jsii.String("asmServer"),
//   		AsmUser: jsii.String("asmUser"),
//   		CharLengthSemantics: jsii.String("charLengthSemantics"),
//   		DirectPathNoLog: jsii.Boolean(false),
//   		DirectPathParallelLoad: jsii.Boolean(false),
//   		EnableHomogenousTablespace: jsii.Boolean(false),
//   		ExtraArchivedLogDestIds: []interface{}{
//   			jsii.Number(123),
//   		},
//   		FailTasksOnLobTruncation: jsii.Boolean(false),
//   		NumberDatatypeScale: jsii.Number(123),
//   		OraclePathPrefix: jsii.String("oraclePathPrefix"),
//   		ParallelAsmReadThreads: jsii.Number(123),
//   		ReadAheadBlocks: jsii.Number(123),
//   		ReadTableSpaceName: jsii.Boolean(false),
//   		ReplacePathPrefix: jsii.Boolean(false),
//   		RetryInterval: jsii.Number(123),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   		SecretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		SecurityDbEncryption: jsii.String("securityDbEncryption"),
//   		SecurityDbEncryptionName: jsii.String("securityDbEncryptionName"),
//   		SpatialDataOptionToGeoJsonFunctionName: jsii.String("spatialDataOptionToGeoJsonFunctionName"),
//   		StandbyDelayTime: jsii.Number(123),
//   		UseAlternateFolderForOnline: jsii.Boolean(false),
//   		UseBFile: jsii.Boolean(false),
//   		UseDirectPathFullLoad: jsii.Boolean(false),
//   		UseLogminerReader: jsii.Boolean(false),
//   		UsePathPrefix: jsii.String("usePathPrefix"),
//   	},
//   	Password: jsii.String("password"),
//   	Port: jsii.Number(123),
//   	PostgreSqlSettings: &PostgreSqlSettingsProperty{
//   		AfterConnectScript: jsii.String("afterConnectScript"),
//   		CaptureDdls: jsii.Boolean(false),
//   		DdlArtifactsSchema: jsii.String("ddlArtifactsSchema"),
//   		ExecuteTimeout: jsii.Number(123),
//   		FailTasksOnLobTruncation: jsii.Boolean(false),
//   		HeartbeatEnable: jsii.Boolean(false),
//   		HeartbeatFrequency: jsii.Number(123),
//   		HeartbeatSchema: jsii.String("heartbeatSchema"),
//   		MapBooleanAsBoolean: jsii.Boolean(false),
//   		MaxFileSize: jsii.Number(123),
//   		PluginName: jsii.String("pluginName"),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		SlotName: jsii.String("slotName"),
//   	},
//   	RedisSettings: &RedisSettingsProperty{
//   		AuthPassword: jsii.String("authPassword"),
//   		AuthType: jsii.String("authType"),
//   		AuthUserName: jsii.String("authUserName"),
//   		Port: jsii.Number(123),
//   		ServerName: jsii.String("serverName"),
//   		SslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		SslSecurityProtocol: jsii.String("sslSecurityProtocol"),
//   	},
//   	RedshiftSettings: &RedshiftSettingsProperty{
//   		AcceptAnyDate: jsii.Boolean(false),
//   		AfterConnectScript: jsii.String("afterConnectScript"),
//   		BucketFolder: jsii.String("bucketFolder"),
//   		BucketName: jsii.String("bucketName"),
//   		CaseSensitiveNames: jsii.Boolean(false),
//   		CompUpdate: jsii.Boolean(false),
//   		ConnectionTimeout: jsii.Number(123),
//   		DateFormat: jsii.String("dateFormat"),
//   		EmptyAsNull: jsii.Boolean(false),
//   		EncryptionMode: jsii.String("encryptionMode"),
//   		ExplicitIds: jsii.Boolean(false),
//   		FileTransferUploadStreams: jsii.Number(123),
//   		LoadTimeout: jsii.Number(123),
//   		MapBooleanAsBoolean: jsii.Boolean(false),
//   		MaxFileSize: jsii.Number(123),
//   		RemoveQuotes: jsii.Boolean(false),
//   		ReplaceChars: jsii.String("replaceChars"),
//   		ReplaceInvalidChars: jsii.String("replaceInvalidChars"),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		ServerSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		TimeFormat: jsii.String("timeFormat"),
//   		TrimBlanks: jsii.Boolean(false),
//   		TruncateColumns: jsii.Boolean(false),
//   		WriteBufferSize: jsii.Number(123),
//   	},
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	S3Settings: &S3SettingsProperty{
//   		AddColumnName: jsii.Boolean(false),
//   		BucketFolder: jsii.String("bucketFolder"),
//   		BucketName: jsii.String("bucketName"),
//   		CannedAclForObjects: jsii.String("cannedAclForObjects"),
//   		CdcInsertsAndUpdates: jsii.Boolean(false),
//   		CdcInsertsOnly: jsii.Boolean(false),
//   		CdcMaxBatchInterval: jsii.Number(123),
//   		CdcMinFileSize: jsii.Number(123),
//   		CdcPath: jsii.String("cdcPath"),
//   		CompressionType: jsii.String("compressionType"),
//   		CsvDelimiter: jsii.String("csvDelimiter"),
//   		CsvNoSupValue: jsii.String("csvNoSupValue"),
//   		CsvNullValue: jsii.String("csvNullValue"),
//   		CsvRowDelimiter: jsii.String("csvRowDelimiter"),
//   		DataFormat: jsii.String("dataFormat"),
//   		DataPageSize: jsii.Number(123),
//   		DatePartitionDelimiter: jsii.String("datePartitionDelimiter"),
//   		DatePartitionEnabled: jsii.Boolean(false),
//   		DatePartitionSequence: jsii.String("datePartitionSequence"),
//   		DatePartitionTimezone: jsii.String("datePartitionTimezone"),
//   		DictPageSizeLimit: jsii.Number(123),
//   		EnableStatistics: jsii.Boolean(false),
//   		EncodingType: jsii.String("encodingType"),
//   		EncryptionMode: jsii.String("encryptionMode"),
//   		ExternalTableDefinition: jsii.String("externalTableDefinition"),
//   		IgnoreHeaderRows: jsii.Number(123),
//   		IncludeOpForFullLoad: jsii.Boolean(false),
//   		MaxFileSize: jsii.Number(123),
//   		ParquetTimestampInMillisecond: jsii.Boolean(false),
//   		ParquetVersion: jsii.String("parquetVersion"),
//   		PreserveTransactions: jsii.Boolean(false),
//   		Rfc4180: jsii.Boolean(false),
//   		RowGroupLength: jsii.Number(123),
//   		ServerSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		ServiceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		TimestampColumnName: jsii.String("timestampColumnName"),
//   		UseCsvNoSupValue: jsii.Boolean(false),
//   		UseTaskStartTimeForFullLoadTimestamp: jsii.Boolean(false),
//   	},
//   	ServerName: jsii.String("serverName"),
//   	SslMode: jsii.String("sslMode"),
//   	SybaseSettings: &SybaseSettingsProperty{
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html
//
type CfnEndpointProps struct {
	// The type of endpoint.
	//
	// Valid values are `source` and `target` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-endpointtype
	//
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The type of engine for the endpoint, depending on the `EndpointType` value.
	//
	// *Valid values* : `mysql` | `oracle` | `postgres` | `mariadb` | `aurora` | `aurora-postgresql` | `opensearch` | `redshift` | `s3` | `db2` | `azuredb` | `sybase` | `dynamodb` | `mongodb` | `kinesis` | `kafka` | `elasticsearch` | `docdb` | `sqlserver` | `neptune`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-enginename
	//
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// The Amazon Resource Name (ARN) for the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, don't specify `DatabaseName` . To migrate to a specific database, use this setting and `targetDbType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Settings in JSON format for the source and target DocumentDB endpoint.
	//
	// For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-docdbsettings
	//
	DocDbSettings interface{} `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-dynamodbsettings
	//
	DynamoDbSettings interface{} `field:"optional" json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-elasticsearchsettings
	//
	ElasticsearchSettings interface{} `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-endpointidentifier
	//
	EndpointIdentifier *string `field:"optional" json:"endpointIdentifier" yaml:"endpointIdentifier"`
	// Additional attributes associated with the connection.
	//
	// Each attribute is specified as a name-value pair associated by an equal sign (=). Multiple attributes are separated by a semicolon (;) with no additional white space. For information on the attributes available for connecting your source or target endpoint, see [Working with AWS DMS Endpoints](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-extraconnectionattributes
	//
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Settings in JSON format for the source GCP MySQL endpoint.
	//
	// These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-gcpmysqlsettings
	//
	GcpMySqlSettings interface{} `field:"optional" json:"gcpMySqlSettings" yaml:"gcpMySqlSettings"`
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-ibmdb2settings
	//
	IbmDb2Settings interface{} `field:"optional" json:"ibmDb2Settings" yaml:"ibmDb2Settings"`
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-kafkasettings
	//
	KafkaSettings interface{} `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-kinesissettings
	//
	KinesisSettings interface{} `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-microsoftsqlserversettings
	//
	MicrosoftSqlServerSettings interface{} `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// Settings in JSON format for the source MongoDB endpoint.
	//
	// For more information about the available settings, see [Using MongoDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-mongodbsettings
	//
	MongoDbSettings interface{} `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// Settings in JSON format for the source and target MySQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-mysqlsettings
	//
	MySqlSettings interface{} `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// Settings in JSON format for the target Amazon Neptune endpoint.
	//
	// For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-neptunesettings
	//
	NeptuneSettings interface{} `field:"optional" json:"neptuneSettings" yaml:"neptuneSettings"`
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-oraclesettings
	//
	OracleSettings interface{} `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// The password to be used to log in to the endpoint database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-postgresqlsettings
	//
	PostgreSqlSettings interface{} `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// Settings in JSON format for the target Redis endpoint.
	//
	// For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-redissettings
	//
	RedisSettings interface{} `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// Settings in JSON format for the Amazon Redshift endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-redshiftsettings
	//
	RedshiftSettings interface{} `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-resourceidentifier
	//
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-s3settings
	//
	S3Settings interface{} `field:"optional" json:"s3Settings" yaml:"s3Settings"`
	// The name of the server where the endpoint database resides.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-servername
	//
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none` .
	//
	// > When `engine_name` is set to S3, the only allowed value is `none` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-sslmode
	//
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-sybasesettings
	//
	SybaseSettings interface{} `field:"optional" json:"sybaseSettings" yaml:"sybaseSettings"`
	// One or more tags to be assigned to the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user name to be used to log in to the endpoint database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

