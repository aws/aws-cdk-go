package previewawsdmsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdms/previewawsdmsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DMS::Endpoint` resource specifies an AWS DMS endpoint.
//
// Currently, AWS CloudFormation supports all AWS DMS endpoint types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnEndpointPropsMixin(&CfnEndpointMixinProps{
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
//   	EndpointType: jsii.String("endpointType"),
//   	EngineName: jsii.String("engineName"),
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html
//
type CfnEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEndpointPropsMixin
type jsiiProxy_CfnEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEndpointPropsMixin) Props() *CfnEndpointMixinProps {
	var returns *CfnEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DMS::Endpoint`.
func NewCfnEndpointPropsMixin(props *CfnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DMS::Endpoint`.
func NewCfnEndpointPropsMixin_Override(c CfnEndpointPropsMixin, props *CfnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_dms.mixins.CfnEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

