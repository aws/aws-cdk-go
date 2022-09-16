package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::DMS::Endpoint`.
//
// The `AWS::DMS::Endpoint` resource specifies an AWS DMS endpoint.
//
// Currently, AWS CloudFormation supports all AWS DMS endpoint types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpoint := awscdk.Aws_dms.NewCfnEndpoint(this, jsii.String("MyCfnEndpoint"), &cfnEndpointProps{
//   	endpointType: jsii.String("endpointType"),
//   	engineName: jsii.String("engineName"),
//
//   	// the properties below are optional
//   	certificateArn: jsii.String("certificateArn"),
//   	databaseName: jsii.String("databaseName"),
//   	docDbSettings: &docDbSettingsProperty{
//   		docsToInvestigate: jsii.Number(123),
//   		extractDocId: jsii.Boolean(false),
//   		nestingLevel: jsii.String("nestingLevel"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	dynamoDbSettings: &dynamoDbSettingsProperty{
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	elasticsearchSettings: &elasticsearchSettingsProperty{
//   		endpointUri: jsii.String("endpointUri"),
//   		errorRetryDuration: jsii.Number(123),
//   		fullLoadErrorPercentage: jsii.Number(123),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	endpointIdentifier: jsii.String("endpointIdentifier"),
//   	extraConnectionAttributes: jsii.String("extraConnectionAttributes"),
//   	gcpMySqlSettings: &gcpMySQLSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		databaseName: jsii.String("databaseName"),
//   		eventsPollInterval: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		parallelLoadThreads: jsii.Number(123),
//   		password: jsii.String("password"),
//   		port: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverName: jsii.String("serverName"),
//   		serverTimezone: jsii.String("serverTimezone"),
//   		username: jsii.String("username"),
//   	},
//   	ibmDb2Settings: &ibmDb2SettingsProperty{
//   		currentLsn: jsii.String("currentLsn"),
//   		maxKBytesPerRead: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		setDataCaptureChanges: jsii.Boolean(false),
//   	},
//   	kafkaSettings: &kafkaSettingsProperty{
//   		broker: jsii.String("broker"),
//   		includeControlDetails: jsii.Boolean(false),
//   		includeNullAndEmpty: jsii.Boolean(false),
//   		includePartitionValue: jsii.Boolean(false),
//   		includeTableAlterOperations: jsii.Boolean(false),
//   		includeTransactionDetails: jsii.Boolean(false),
//   		messageFormat: jsii.String("messageFormat"),
//   		messageMaxBytes: jsii.Number(123),
//   		noHexPrefix: jsii.Boolean(false),
//   		partitionIncludeSchemaTable: jsii.Boolean(false),
//   		saslPassword: jsii.String("saslPassword"),
//   		saslUserName: jsii.String("saslUserName"),
//   		securityProtocol: jsii.String("securityProtocol"),
//   		sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		sslClientCertificateArn: jsii.String("sslClientCertificateArn"),
//   		sslClientKeyArn: jsii.String("sslClientKeyArn"),
//   		sslClientKeyPassword: jsii.String("sslClientKeyPassword"),
//   		topic: jsii.String("topic"),
//   	},
//   	kinesisSettings: &kinesisSettingsProperty{
//   		includeControlDetails: jsii.Boolean(false),
//   		includeNullAndEmpty: jsii.Boolean(false),
//   		includePartitionValue: jsii.Boolean(false),
//   		includeTableAlterOperations: jsii.Boolean(false),
//   		includeTransactionDetails: jsii.Boolean(false),
//   		messageFormat: jsii.String("messageFormat"),
//   		noHexPrefix: jsii.Boolean(false),
//   		partitionIncludeSchemaTable: jsii.Boolean(false),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	microsoftSqlServerSettings: &microsoftSqlServerSettingsProperty{
//   		bcpPacketSize: jsii.Number(123),
//   		controlTablesFileGroup: jsii.String("controlTablesFileGroup"),
//   		querySingleAlwaysOnNode: jsii.Boolean(false),
//   		readBackupOnly: jsii.Boolean(false),
//   		safeguardPolicy: jsii.String("safeguardPolicy"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		useBcpFullLoad: jsii.Boolean(false),
//   		useThirdPartyBackupDevice: jsii.Boolean(false),
//   	},
//   	mongoDbSettings: &mongoDbSettingsProperty{
//   		authMechanism: jsii.String("authMechanism"),
//   		authSource: jsii.String("authSource"),
//   		authType: jsii.String("authType"),
//   		databaseName: jsii.String("databaseName"),
//   		docsToInvestigate: jsii.String("docsToInvestigate"),
//   		extractDocId: jsii.String("extractDocId"),
//   		nestingLevel: jsii.String("nestingLevel"),
//   		password: jsii.String("password"),
//   		port: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverName: jsii.String("serverName"),
//   		username: jsii.String("username"),
//   	},
//   	mySqlSettings: &mySqlSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		cleanSourceMetadataOnMismatch: jsii.Boolean(false),
//   		eventsPollInterval: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		parallelLoadThreads: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverTimezone: jsii.String("serverTimezone"),
//   		targetDbType: jsii.String("targetDbType"),
//   	},
//   	neptuneSettings: &neptuneSettingsProperty{
//   		errorRetryDuration: jsii.Number(123),
//   		iamAuthEnabled: jsii.Boolean(false),
//   		maxFileSize: jsii.Number(123),
//   		maxRetryCount: jsii.Number(123),
//   		s3BucketFolder: jsii.String("s3BucketFolder"),
//   		s3BucketName: jsii.String("s3BucketName"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   	},
//   	oracleSettings: &oracleSettingsProperty{
//   		accessAlternateDirectly: jsii.Boolean(false),
//   		additionalArchivedLogDestId: jsii.Number(123),
//   		addSupplementalLogging: jsii.Boolean(false),
//   		allowSelectNestedTables: jsii.Boolean(false),
//   		archivedLogDestId: jsii.Number(123),
//   		archivedLogsOnly: jsii.Boolean(false),
//   		asmPassword: jsii.String("asmPassword"),
//   		asmServer: jsii.String("asmServer"),
//   		asmUser: jsii.String("asmUser"),
//   		charLengthSemantics: jsii.String("charLengthSemantics"),
//   		directPathNoLog: jsii.Boolean(false),
//   		directPathParallelLoad: jsii.Boolean(false),
//   		enableHomogenousTablespace: jsii.Boolean(false),
//   		extraArchivedLogDestIds: []interface{}{
//   			jsii.Number(123),
//   		},
//   		failTasksOnLobTruncation: jsii.Boolean(false),
//   		numberDatatypeScale: jsii.Number(123),
//   		oraclePathPrefix: jsii.String("oraclePathPrefix"),
//   		parallelAsmReadThreads: jsii.Number(123),
//   		readAheadBlocks: jsii.Number(123),
//   		readTableSpaceName: jsii.Boolean(false),
//   		replacePathPrefix: jsii.Boolean(false),
//   		retryInterval: jsii.Number(123),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerOracleAsmAccessRoleArn: jsii.String("secretsManagerOracleAsmAccessRoleArn"),
//   		secretsManagerOracleAsmSecretId: jsii.String("secretsManagerOracleAsmSecretId"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		securityDbEncryption: jsii.String("securityDbEncryption"),
//   		securityDbEncryptionName: jsii.String("securityDbEncryptionName"),
//   		spatialDataOptionToGeoJsonFunctionName: jsii.String("spatialDataOptionToGeoJsonFunctionName"),
//   		standbyDelayTime: jsii.Number(123),
//   		useAlternateFolderForOnline: jsii.Boolean(false),
//   		useBFile: jsii.Boolean(false),
//   		useDirectPathFullLoad: jsii.Boolean(false),
//   		useLogminerReader: jsii.Boolean(false),
//   		usePathPrefix: jsii.String("usePathPrefix"),
//   	},
//   	password: jsii.String("password"),
//   	port: jsii.Number(123),
//   	postgreSqlSettings: &postgreSqlSettingsProperty{
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		captureDdls: jsii.Boolean(false),
//   		ddlArtifactsSchema: jsii.String("ddlArtifactsSchema"),
//   		executeTimeout: jsii.Number(123),
//   		failTasksOnLobTruncation: jsii.Boolean(false),
//   		heartbeatEnable: jsii.Boolean(false),
//   		heartbeatFrequency: jsii.Number(123),
//   		heartbeatSchema: jsii.String("heartbeatSchema"),
//   		maxFileSize: jsii.Number(123),
//   		pluginName: jsii.String("pluginName"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		slotName: jsii.String("slotName"),
//   	},
//   	redisSettings: &redisSettingsProperty{
//   		authPassword: jsii.String("authPassword"),
//   		authType: jsii.String("authType"),
//   		authUserName: jsii.String("authUserName"),
//   		port: jsii.Number(123),
//   		serverName: jsii.String("serverName"),
//   		sslCaCertificateArn: jsii.String("sslCaCertificateArn"),
//   		sslSecurityProtocol: jsii.String("sslSecurityProtocol"),
//   	},
//   	redshiftSettings: &redshiftSettingsProperty{
//   		acceptAnyDate: jsii.Boolean(false),
//   		afterConnectScript: jsii.String("afterConnectScript"),
//   		bucketFolder: jsii.String("bucketFolder"),
//   		bucketName: jsii.String("bucketName"),
//   		caseSensitiveNames: jsii.Boolean(false),
//   		compUpdate: jsii.Boolean(false),
//   		connectionTimeout: jsii.Number(123),
//   		dateFormat: jsii.String("dateFormat"),
//   		emptyAsNull: jsii.Boolean(false),
//   		encryptionMode: jsii.String("encryptionMode"),
//   		explicitIds: jsii.Boolean(false),
//   		fileTransferUploadStreams: jsii.Number(123),
//   		loadTimeout: jsii.Number(123),
//   		maxFileSize: jsii.Number(123),
//   		removeQuotes: jsii.Boolean(false),
//   		replaceChars: jsii.String("replaceChars"),
//   		replaceInvalidChars: jsii.String("replaceInvalidChars"),
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		timeFormat: jsii.String("timeFormat"),
//   		trimBlanks: jsii.Boolean(false),
//   		truncateColumns: jsii.Boolean(false),
//   		writeBufferSize: jsii.Number(123),
//   	},
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	s3Settings: &s3SettingsProperty{
//   		addColumnName: jsii.Boolean(false),
//   		bucketFolder: jsii.String("bucketFolder"),
//   		bucketName: jsii.String("bucketName"),
//   		cannedAclForObjects: jsii.String("cannedAclForObjects"),
//   		cdcInsertsAndUpdates: jsii.Boolean(false),
//   		cdcInsertsOnly: jsii.Boolean(false),
//   		cdcMaxBatchInterval: jsii.Number(123),
//   		cdcMinFileSize: jsii.Number(123),
//   		cdcPath: jsii.String("cdcPath"),
//   		compressionType: jsii.String("compressionType"),
//   		csvDelimiter: jsii.String("csvDelimiter"),
//   		csvNoSupValue: jsii.String("csvNoSupValue"),
//   		csvNullValue: jsii.String("csvNullValue"),
//   		csvRowDelimiter: jsii.String("csvRowDelimiter"),
//   		dataFormat: jsii.String("dataFormat"),
//   		dataPageSize: jsii.Number(123),
//   		datePartitionDelimiter: jsii.String("datePartitionDelimiter"),
//   		datePartitionEnabled: jsii.Boolean(false),
//   		datePartitionSequence: jsii.String("datePartitionSequence"),
//   		datePartitionTimezone: jsii.String("datePartitionTimezone"),
//   		dictPageSizeLimit: jsii.Number(123),
//   		enableStatistics: jsii.Boolean(false),
//   		encodingType: jsii.String("encodingType"),
//   		encryptionMode: jsii.String("encryptionMode"),
//   		externalTableDefinition: jsii.String("externalTableDefinition"),
//   		ignoreHeaderRows: jsii.Number(123),
//   		includeOpForFullLoad: jsii.Boolean(false),
//   		maxFileSize: jsii.Number(123),
//   		parquetTimestampInMillisecond: jsii.Boolean(false),
//   		parquetVersion: jsii.String("parquetVersion"),
//   		preserveTransactions: jsii.Boolean(false),
//   		rfc4180: jsii.Boolean(false),
//   		rowGroupLength: jsii.Number(123),
//   		serverSideEncryptionKmsKeyId: jsii.String("serverSideEncryptionKmsKeyId"),
//   		serviceAccessRoleArn: jsii.String("serviceAccessRoleArn"),
//   		timestampColumnName: jsii.String("timestampColumnName"),
//   		useCsvNoSupValue: jsii.Boolean(false),
//   		useTaskStartTimeForFullLoadTimestamp: jsii.Boolean(false),
//   	},
//   	serverName: jsii.String("serverName"),
//   	sslMode: jsii.String("sslMode"),
//   	sybaseSettings: &sybaseSettingsProperty{
//   		secretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		secretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	username: jsii.String("username"),
//   })
//
type CfnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A value that can be used for cross-account validation.
	AttrExternalId() *string
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn() *string
	SetCertificateArn(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, don't specify `DatabaseName` . To migrate to a specific database, use this setting and `targetDbType` .
	DatabaseName() *string
	SetDatabaseName(val *string)
	// Settings in JSON format for the source and target DocumentDB endpoint.
	//
	// For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
	DocDbSettings() interface{}
	SetDocDbSettings(val interface{})
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	DynamoDbSettings() interface{}
	SetDynamoDbSettings(val interface{})
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
	ElasticsearchSettings() interface{}
	SetElasticsearchSettings(val interface{})
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	EndpointIdentifier() *string
	SetEndpointIdentifier(val *string)
	// The type of endpoint.
	//
	// Valid values are `source` and `target` .
	EndpointType() *string
	SetEndpointType(val *string)
	// The type of engine for the endpoint, depending on the `EndpointType` value.
	//
	// *Valid values* : `mysql` | `oracle` | `postgres` | `mariadb` | `aurora` | `aurora-postgresql` | `opensearch` | `redshift` | `s3` | `db2` | `azuredb` | `sybase` | `dynamodb` | `mongodb` | `kinesis` | `kafka` | `elasticsearch` | `docdb` | `sqlserver` | `neptune`.
	EngineName() *string
	SetEngineName(val *string)
	// Additional attributes associated with the connection.
	//
	// Each attribute is specified as a name-value pair associated by an equal sign (=). Multiple attributes are separated by a semicolon (;) with no additional white space. For information on the attributes available for connecting your source or target endpoint, see [Working with AWS DMS Endpoints](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the *AWS Database Migration Service User Guide* .
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	// Settings in JSON format for the source GCP MySQL endpoint.
	//
	// These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	GcpMySqlSettings() interface{}
	SetGcpMySqlSettings(val interface{})
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	IbmDb2Settings() interface{}
	SetIbmDb2Settings(val interface{})
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KafkaSettings() interface{}
	SetKafkaSettings(val interface{})
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KinesisSettings() interface{}
	SetKinesisSettings(val interface{})
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MicrosoftSqlServerSettings() interface{}
	SetMicrosoftSqlServerSettings(val interface{})
	// Settings in JSON format for the source MongoDB endpoint.
	//
	// For more information about the available settings, see [Using MongoDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
	MongoDbSettings() interface{}
	SetMongoDbSettings(val interface{})
	// Settings in JSON format for the source and target MySQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MySqlSettings() interface{}
	SetMySqlSettings(val interface{})
	// Settings in JSON format for the target Amazon Neptune endpoint.
	//
	// For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	NeptuneSettings() interface{}
	SetNeptuneSettings(val interface{})
	// The tree node.
	Node() constructs.Node
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	OracleSettings() interface{}
	SetOracleSettings(val interface{})
	// The password to be used to log in to the endpoint database.
	Password() *string
	SetPassword(val *string)
	// The port used by the endpoint database.
	Port() *float64
	SetPort(val *float64)
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	PostgreSqlSettings() interface{}
	SetPostgreSqlSettings(val interface{})
	// Settings in JSON format for the target Redis endpoint.
	//
	// For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	RedisSettings() interface{}
	SetRedisSettings(val interface{})
	// Settings in JSON format for the Amazon Redshift endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	RedshiftSettings() interface{}
	SetRedshiftSettings(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
	S3Settings() interface{}
	SetS3Settings(val interface{})
	// The name of the server where the endpoint database resides.
	ServerName() *string
	SetServerName(val *string)
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none` .
	//
	// > When `engine_name` is set to S3, the only allowed value is `none` .
	SslMode() *string
	SetSslMode(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	SybaseSettings() interface{}
	SetSybaseSettings(val interface{})
	// One or more tags to be assigned to the endpoint.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// The user name to be used to log in to the endpoint database.
	Username() *string
	SetUsername(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEndpoint
type jsiiProxy_CfnEndpoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpoint) AttrExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DocDbSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"docDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) DynamoDbSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ElasticsearchSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EndpointIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) GcpMySqlSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpMySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) IbmDb2Settings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ibmDb2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) KafkaSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) KinesisSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) MicrosoftSqlServerSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSqlServerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) MongoDbSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) MySqlSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) NeptuneSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) OracleSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) PostgreSqlSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) RedisSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) RedshiftSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) S3Settings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) SybaseSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sybaseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


// Create a new `AWS::DMS::Endpoint`.
func NewCfnEndpoint(scope constructs.Construct, id *string, props *CfnEndpointProps) CfnEndpoint {
	_init_.Initialize()

	if err := validateNewCfnEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DMS::Endpoint`.
func NewCfnEndpoint_Override(c CfnEndpoint, scope constructs.Construct, id *string, props *CfnEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetDocDbSettings(val interface{}) {
	if err := j.validateSetDocDbSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"docDbSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetDynamoDbSettings(val interface{}) {
	if err := j.validateSetDynamoDbSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamoDbSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetElasticsearchSettings(val interface{}) {
	if err := j.validateSetElasticsearchSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetEndpointIdentifier(val *string) {
	_jsii_.Set(
		j,
		"endpointIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetEngineName(val *string) {
	if err := j.validateSetEngineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetExtraConnectionAttributes(val *string) {
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetGcpMySqlSettings(val interface{}) {
	if err := j.validateSetGcpMySqlSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpMySqlSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetIbmDb2Settings(val interface{}) {
	if err := j.validateSetIbmDb2SettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ibmDb2Settings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetKafkaSettings(val interface{}) {
	if err := j.validateSetKafkaSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetKinesisSettings(val interface{}) {
	if err := j.validateSetKinesisSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kinesisSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetMicrosoftSqlServerSettings(val interface{}) {
	if err := j.validateSetMicrosoftSqlServerSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftSqlServerSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetMongoDbSettings(val interface{}) {
	if err := j.validateSetMongoDbSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongoDbSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetMySqlSettings(val interface{}) {
	if err := j.validateSetMySqlSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mySqlSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetNeptuneSettings(val interface{}) {
	if err := j.validateSetNeptuneSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"neptuneSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetOracleSettings(val interface{}) {
	if err := j.validateSetOracleSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oracleSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetPostgreSqlSettings(val interface{}) {
	if err := j.validateSetPostgreSqlSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgreSqlSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetRedisSettings(val interface{}) {
	if err := j.validateSetRedisSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetRedshiftSettings(val interface{}) {
	if err := j.validateSetRedshiftSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redshiftSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetResourceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetS3Settings(val interface{}) {
	if err := j.validateSetS3SettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Settings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetSslMode(val *string) {
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetSybaseSettings(val interface{}) {
	if err := j.validateSetSybaseSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sybaseSettings",
		val,
	)
}

func (j *jsiiProxy_CfnEndpoint)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEndpoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpoint_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEndpoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnEndpoint_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpoint) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEndpoint) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

