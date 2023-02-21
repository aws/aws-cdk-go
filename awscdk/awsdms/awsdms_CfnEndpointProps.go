package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &cfnEndpointProps{
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
//   }
//
type CfnEndpointProps struct {
	// The type of endpoint.
	//
	// Valid values are `source` and `target` .
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The type of engine for the endpoint, depending on the `EndpointType` value.
	//
	// *Valid values* : `mysql` | `oracle` | `postgres` | `mariadb` | `aurora` | `aurora-postgresql` | `opensearch` | `redshift` | `s3` | `db2` | `azuredb` | `sybase` | `dynamodb` | `mongodb` | `kinesis` | `kafka` | `elasticsearch` | `docdb` | `sqlserver` | `neptune`.
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// The Amazon Resource Name (ARN) for the certificate.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The name of the endpoint database.
	//
	// For a MySQL source or target endpoint, don't specify `DatabaseName` . To migrate to a specific database, use this setting and `targetDbType` .
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Settings in JSON format for the source and target DocumentDB endpoint.
	//
	// For more information about other available settings, see [Using extra connections attributes with Amazon DocumentDB as a source](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DocumentDB.html#CHAP_Source.DocumentDB.ECAs) and [Using Amazon DocumentDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DocumentDB.html) in the *AWS Database Migration Service User Guide* .
	DocDbSettings interface{} `field:"optional" json:"docDbSettings" yaml:"docDbSettings"`
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	//
	// For information about other available settings, see [Using object mapping to migrate data to DynamoDB](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.DynamoDB.html#CHAP_Target.DynamoDB.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	DynamoDbSettings interface{} `field:"optional" json:"dynamoDbSettings" yaml:"dynamoDbSettings"`
	// Settings in JSON format for the target OpenSearch endpoint.
	//
	// For more information about the available settings, see [Extra connection attributes when using OpenSearch as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Elasticsearch.html#CHAP_Target.Elasticsearch.Configuration) in the *AWS Database Migration Service User Guide* .
	ElasticsearchSettings interface{} `field:"optional" json:"elasticsearchSettings" yaml:"elasticsearchSettings"`
	// The database endpoint identifier.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	EndpointIdentifier *string `field:"optional" json:"endpointIdentifier" yaml:"endpointIdentifier"`
	// Additional attributes associated with the connection.
	//
	// Each attribute is specified as a name-value pair associated by an equal sign (=). Multiple attributes are separated by a semicolon (;) with no additional white space. For information on the attributes available for connecting your source or target endpoint, see [Working with AWS DMS Endpoints](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Endpoints.html) in the *AWS Database Migration Service User Guide* .
	ExtraConnectionAttributes *string `field:"optional" json:"extraConnectionAttributes" yaml:"extraConnectionAttributes"`
	// Settings in JSON format for the source GCP MySQL endpoint.
	//
	// These settings are much the same as the settings for any MySQL-compatible endpoint. For more information, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	GcpMySqlSettings interface{} `field:"optional" json:"gcpMySqlSettings" yaml:"gcpMySqlSettings"`
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Db2 LUW as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.DB2.html#CHAP_Source.DB2.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	IbmDb2Settings interface{} `field:"optional" json:"ibmDb2Settings" yaml:"ibmDb2Settings"`
	// Settings in JSON format for the target Apache Kafka endpoint.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kafka topic](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kafka.html#CHAP_Target.Kafka.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KafkaSettings interface{} `field:"optional" json:"kafkaSettings" yaml:"kafkaSettings"`
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	//
	// For more information about other available settings, see [Using object mapping to migrate data to a Kinesis data stream](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Kinesis.html#CHAP_Target.Kinesis.ObjectMapping) in the *AWS Database Migration Service User Guide* .
	KinesisSettings interface{} `field:"optional" json:"kinesisSettings" yaml:"kinesisSettings"`
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Region .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Settings in JSON format for the source and target Microsoft SQL Server endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SQL Server as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html#CHAP_Source.SQLServer.ConnectionAttrib) and [Extra connection attributes when using SQL Server as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SQLServer.html#CHAP_Target.SQLServer.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MicrosoftSqlServerSettings interface{} `field:"optional" json:"microsoftSqlServerSettings" yaml:"microsoftSqlServerSettings"`
	// Settings in JSON format for the source MongoDB endpoint.
	//
	// For more information about the available settings, see [Using MongoDB as a target for AWS Database Migration Service](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MongoDB.html#CHAP_Source.MongoDB.Configuration) in the *AWS Database Migration Service User Guide* .
	MongoDbSettings interface{} `field:"optional" json:"mongoDbSettings" yaml:"mongoDbSettings"`
	// Settings in JSON format for the source and target MySQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using MySQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.MySQL.html#CHAP_Source.MySQL.ConnectionAttrib) and [Extra connection attributes when using a MySQL-compatible database as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.MySQL.html#CHAP_Target.MySQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	MySqlSettings interface{} `field:"optional" json:"mySqlSettings" yaml:"mySqlSettings"`
	// Settings in JSON format for the target Amazon Neptune endpoint.
	//
	// For more information about the available settings, see [Specifying endpoint settings for Amazon Neptune as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Neptune.html#CHAP_Target.Neptune.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	NeptuneSettings interface{} `field:"optional" json:"neptuneSettings" yaml:"neptuneSettings"`
	// Settings in JSON format for the source and target Oracle endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using Oracle as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.Oracle.html#CHAP_Source.Oracle.ConnectionAttrib) and [Extra connection attributes when using Oracle as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Oracle.html#CHAP_Target.Oracle.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	OracleSettings interface{} `field:"optional" json:"oracleSettings" yaml:"oracleSettings"`
	// The password to be used to log in to the endpoint database.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port used by the endpoint database.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using PostgreSQL as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib) and [Extra connection attributes when using PostgreSQL as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.PostgreSQL.html#CHAP_Target.PostgreSQL.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	PostgreSqlSettings interface{} `field:"optional" json:"postgreSqlSettings" yaml:"postgreSqlSettings"`
	// Settings in JSON format for the target Redis endpoint.
	//
	// For information about other available settings, see [Specifying endpoint settings for Redis as a target](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redis.html#CHAP_Target.Redis.EndpointSettings) in the *AWS Database Migration Service User Guide* .
	RedisSettings interface{} `field:"optional" json:"redisSettings" yaml:"redisSettings"`
	// Settings in JSON format for the Amazon Redshift endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon Redshift as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.Redshift.html#CHAP_Target.Redshift.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	RedshiftSettings interface{} `field:"optional" json:"redshiftSettings" yaml:"redshiftSettings"`
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	//
	// The value for this parameter can have up to 31 characters. It can contain only ASCII letters, digits, and hyphen ('-'). Also, it can't end with a hyphen or contain two consecutive hyphens, and can only begin with a letter, such as `Example-App-ARN1` .
	//
	// For example, this value might result in the `EndpointArn` value `arn:aws:dms:eu-west-1:012345678901:rep:Example-App-ARN1` . If you don't specify a `ResourceIdentifier` value, AWS DMS generates a default identifier value for the end of `EndpointArn` .
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	//
	// For more information about other available settings, see [Extra connection attributes when using Amazon S3 as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.S3.html#CHAP_Source.S3.Configuring) and [Extra connection attributes when using Amazon S3 as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.S3.html#CHAP_Target.S3.Configuring) in the *AWS Database Migration Service User Guide* .
	S3Settings interface{} `field:"optional" json:"s3Settings" yaml:"s3Settings"`
	// The name of the server where the endpoint database resides.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection. The default is `none` .
	//
	// > When `engine_name` is set to S3, the only allowed value is `none` .
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
	// Settings in JSON format for the source and target SAP ASE endpoint.
	//
	// For information about other available settings, see [Extra connection attributes when using SAP ASE as a source for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SAP.html#CHAP_Source.SAP.ConnectionAttrib) and [Extra connection attributes when using SAP ASE as a target for AWS DMS](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Target.SAP.html#CHAP_Target.SAP.ConnectionAttrib) in the *AWS Database Migration Service User Guide* .
	SybaseSettings interface{} `field:"optional" json:"sybaseSettings" yaml:"sybaseSettings"`
	// One or more tags to be assigned to the endpoint.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The user name to be used to log in to the endpoint database.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

