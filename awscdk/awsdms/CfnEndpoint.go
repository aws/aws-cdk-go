package awsdms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdms"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DMS::Endpoint` resource specifies an AWS DMS endpoint.
//
// Currently, AWS CloudFormation supports all AWS DMS endpoint types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpoint := awscdk.Aws_dms.NewCfnEndpoint(this, jsii.String("MyCfnEndpoint"), &CfnEndpointProps{
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
//   		KeepCsvFiles: jsii.Boolean(false),
//   		LoadTimeout: jsii.Number(123),
//   		MaxFileSize: jsii.Number(123),
//   		MaxKBytesPerRead: jsii.Number(123),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		SetDataCaptureChanges: jsii.Boolean(false),
//   		WriteBufferSize: jsii.Number(123),
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
//   		DatabaseName: jsii.String("databaseName"),
//   		ForceLobLookup: jsii.Boolean(false),
//   		Password: jsii.String("password"),
//   		Port: jsii.Number(123),
//   		QuerySingleAlwaysOnNode: jsii.Boolean(false),
//   		ReadBackupOnly: jsii.Boolean(false),
//   		SafeguardPolicy: jsii.String("safeguardPolicy"),
//   		SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   		SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		ServerName: jsii.String("serverName"),
//   		TlogAccessMode: jsii.String("tlogAccessMode"),
//   		TrimSpaceInChar: jsii.Boolean(false),
//   		UseBcpFullLoad: jsii.Boolean(false),
//   		Username: jsii.String("username"),
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
//   		BabelfishDatabaseName: jsii.String("babelfishDatabaseName"),
//   		CaptureDdls: jsii.Boolean(false),
//   		DatabaseMode: jsii.String("databaseMode"),
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
//   		AddTrailingPaddingCharacter: jsii.Boolean(false),
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
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		ExternalTableDefinition: jsii.String("externalTableDefinition"),
//   		GlueCatalogGeneration: jsii.Boolean(false),
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html
//
type CfnEndpoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsdms.IEndpointRef
	awscdk.ITaggable
	// A value that can be used for cross-account validation.
	AttrExternalId() *string
	AttrId() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	// Settings in JSON format for the source and target DocumentDB endpoint.
	DocDbSettings() interface{}
	SetDocDbSettings(val interface{})
	// Settings in JSON format for the target Amazon DynamoDB endpoint.
	DynamoDbSettings() interface{}
	SetDynamoDbSettings(val interface{})
	// Settings in JSON format for the target OpenSearch endpoint.
	ElasticsearchSettings() interface{}
	SetElasticsearchSettings(val interface{})
	// The database endpoint identifier.
	EndpointIdentifier() *string
	SetEndpointIdentifier(val *string)
	// A reference to a Endpoint resource.
	EndpointRef() *interfacesawsdms.EndpointReference
	// The type of endpoint.
	EndpointType() *string
	SetEndpointType(val *string)
	// The type of engine for the endpoint, depending on the `EndpointType` value.
	EngineName() *string
	SetEngineName(val *string)
	Env() *interfaces.ResourceEnvironment
	// Additional attributes associated with the connection.
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	// Settings in JSON format for the source GCP MySQL endpoint.
	GcpMySqlSettings() interface{}
	SetGcpMySqlSettings(val interface{})
	// Settings in JSON format for the source IBM Db2 LUW endpoint.
	IbmDb2Settings() interface{}
	SetIbmDb2Settings(val interface{})
	// Settings in JSON format for the target Apache Kafka endpoint.
	KafkaSettings() interface{}
	SetKafkaSettings(val interface{})
	// Settings in JSON format for the target endpoint for Amazon Kinesis Data Streams.
	KinesisSettings() interface{}
	SetKinesisSettings(val interface{})
	// An AWS KMS key identifier that is used to encrypt the connection parameters for the endpoint.
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
	MicrosoftSqlServerSettings() interface{}
	SetMicrosoftSqlServerSettings(val interface{})
	// Settings in JSON format for the source MongoDB endpoint.
	MongoDbSettings() interface{}
	SetMongoDbSettings(val interface{})
	// Settings in JSON format for the source and target MySQL endpoint.
	MySqlSettings() interface{}
	SetMySqlSettings(val interface{})
	// Settings in JSON format for the target Amazon Neptune endpoint.
	NeptuneSettings() interface{}
	SetNeptuneSettings(val interface{})
	// The tree node.
	Node() constructs.Node
	// Settings in JSON format for the source and target Oracle endpoint.
	OracleSettings() interface{}
	SetOracleSettings(val interface{})
	// The password to be used to log in to the endpoint database.
	Password() *string
	SetPassword(val *string)
	// The port used by the endpoint database.
	Port() *float64
	SetPort(val *float64)
	// Settings in JSON format for the source and target PostgreSQL endpoint.
	PostgreSqlSettings() interface{}
	SetPostgreSqlSettings(val interface{})
	// Settings in JSON format for the target Redis endpoint.
	RedisSettings() interface{}
	SetRedisSettings(val interface{})
	// Settings in JSON format for the Amazon Redshift endpoint.
	RedshiftSettings() interface{}
	SetRedshiftSettings(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A display name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	// Settings in JSON format for the source and target Amazon S3 endpoint.
	S3Settings() interface{}
	SetS3Settings(val interface{})
	// The name of the server where the endpoint database resides.
	ServerName() *string
	SetServerName(val *string)
	// The Secure Sockets Layer (SSL) mode to use for the SSL connection.
	//
	// The default is `none` .
	SslMode() *string
	SetSslMode(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Settings in JSON format for the source and target SAP ASE endpoint.
	SybaseSettings() interface{}
	SetSybaseSettings(val interface{})
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// One or more tags to be assigned to the endpoint.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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
	internal.Type__interfacesawsdmsIEndpointRef
	internal.Type__awscdkITaggable
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

func (j *jsiiProxy_CfnEndpoint) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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

func (j *jsiiProxy_CfnEndpoint) EndpointRef() *interfacesawsdms.EndpointReference {
	var returns *interfacesawsdms.EndpointReference
	_jsii_.Get(
		j,
		"endpointRef",
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

func (j *jsiiProxy_CfnEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
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

func (j *jsiiProxy_CfnEndpoint) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
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

func (j *jsiiProxy_CfnEndpoint)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
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

// Creates a new IEndpointRef from a endpointId.
func CfnEndpoint_FromEndpointId(scope constructs.Construct, id *string, endpointId *string) interfacesawsdms.IEndpointRef {
	_init_.Initialize()

	if err := validateCfnEndpoint_FromEndpointIdParameters(scope, id, endpointId); err != nil {
		panic(err)
	}
	var returns interfacesawsdms.IEndpointRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"fromEndpointId",
		[]interface{}{scope, id, endpointId},
		&returns,
	)

	return returns
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

// Check whether the given object is a CfnResource.
func CfnEndpoint_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpoint_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dms.CfnEndpoint",
		"isCfnResource",
		[]interface{}{x},
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

func (c *jsiiProxy_CfnEndpoint) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnEndpoint) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
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

func (c *jsiiProxy_CfnEndpoint) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpoint) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CfnEndpoint) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
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

func (c *jsiiProxy_CfnEndpoint) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

