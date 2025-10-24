package awskinesisfirehose

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_kinesisfirehose.BackupMode",
		reflect.TypeOf((*BackupMode)(nil)).Elem(),
		map[string]interface{}{
			"ALL": BackupMode_ALL,
			"FAILED": BackupMode_FAILED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		reflect.TypeOf((*CfnDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonOpenSearchServerlessDestinationConfiguration", GoGetter: "AmazonOpenSearchServerlessDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "amazonopensearchserviceDestinationConfiguration", GoGetter: "AmazonopensearchserviceDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "databaseSourceConfiguration", GoGetter: "DatabaseSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamEncryptionConfigurationInput", GoGetter: "DeliveryStreamEncryptionConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamRef", GoGetter: "DeliveryStreamRef"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamType", GoGetter: "DeliveryStreamType"},
			_jsii_.MemberProperty{JsiiProperty: "directPutSourceConfiguration", GoGetter: "DirectPutSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "elasticsearchDestinationConfiguration", GoGetter: "ElasticsearchDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "extendedS3DestinationConfiguration", GoGetter: "ExtendedS3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "httpEndpointDestinationConfiguration", GoGetter: "HttpEndpointDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "icebergDestinationConfiguration", GoGetter: "IcebergDestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamSourceConfiguration", GoGetter: "KinesisStreamSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mskSourceConfiguration", GoGetter: "MskSourceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "redshiftDestinationConfiguration", GoGetter: "RedshiftDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "s3DestinationConfiguration", GoGetter: "S3DestinationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "snowflakeDestinationConfiguration", GoGetter: "SnowflakeDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "splunkDestinationConfiguration", GoGetter: "SplunkDestinationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeliveryStreamRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AmazonOpenSearchServerlessBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonOpenSearchServerlessBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AmazonOpenSearchServerlessRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonOpenSearchServerlessRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AmazonopensearchserviceBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonopensearchserviceBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AmazonopensearchserviceDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AmazonopensearchserviceRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_AmazonopensearchserviceRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.BufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_BufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.CatalogConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_CatalogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.CloudWatchLoggingOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_CloudWatchLoggingOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.CopyCommandProperty",
		reflect.TypeOf((*CfnDeliveryStream_CopyCommandProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DataFormatConversionConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DataFormatConversionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DatabaseColumnsProperty",
		reflect.TypeOf((*CfnDeliveryStream_DatabaseColumnsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DatabaseSourceAuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DatabaseSourceAuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DatabaseSourceConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DatabaseSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DatabaseSourceVPCConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DatabaseSourceVPCConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DatabaseTablesProperty",
		reflect.TypeOf((*CfnDeliveryStream_DatabaseTablesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DatabasesProperty",
		reflect.TypeOf((*CfnDeliveryStream_DatabasesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DeliveryStreamEncryptionConfigurationInputProperty",
		reflect.TypeOf((*CfnDeliveryStream_DeliveryStreamEncryptionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DeserializerProperty",
		reflect.TypeOf((*CfnDeliveryStream_DeserializerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DestinationTableConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DestinationTableConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DirectPutSourceConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DirectPutSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DocumentIdOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_DocumentIdOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.DynamicPartitioningConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_DynamicPartitioningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ElasticsearchBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_ElasticsearchBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ElasticsearchDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ElasticsearchRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_ElasticsearchRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ExtendedS3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.HiveJsonSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_HiveJsonSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointCommonAttributeProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointCommonAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.HttpEndpointRequestConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_HttpEndpointRequestConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.IcebergDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_IcebergDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.InputFormatConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_InputFormatConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.KMSEncryptionConfigProperty",
		reflect.TypeOf((*CfnDeliveryStream_KMSEncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.KinesisStreamSourceConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_KinesisStreamSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.MSKSourceConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_MSKSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.OpenXJsonSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_OpenXJsonSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.OrcSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_OrcSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.OutputFormatConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_OutputFormatConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ParquetSerDeProperty",
		reflect.TypeOf((*CfnDeliveryStream_ParquetSerDeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.PartitionFieldProperty",
		reflect.TypeOf((*CfnDeliveryStream_PartitionFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.PartitionSpecProperty",
		reflect.TypeOf((*CfnDeliveryStream_PartitionSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ProcessingConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_ProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ProcessorParameterProperty",
		reflect.TypeOf((*CfnDeliveryStream_ProcessorParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.ProcessorProperty",
		reflect.TypeOf((*CfnDeliveryStream_ProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.RedshiftDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_RedshiftDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.RedshiftRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_RedshiftRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.RetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_RetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.S3DestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_S3DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SchemaConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SchemaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SchemaEvolutionConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SchemaEvolutionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SecretsManagerConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SecretsManagerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SerializerProperty",
		reflect.TypeOf((*CfnDeliveryStream_SerializerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SnowflakeBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_SnowflakeBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SnowflakeDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SnowflakeDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SnowflakeRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_SnowflakeRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SnowflakeRoleConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SnowflakeRoleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SnowflakeVpcConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SnowflakeVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SplunkBufferingHintsProperty",
		reflect.TypeOf((*CfnDeliveryStream_SplunkBufferingHintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SplunkDestinationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_SplunkDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.SplunkRetryOptionsProperty",
		reflect.TypeOf((*CfnDeliveryStream_SplunkRetryOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.TableCreationConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_TableCreationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream.VpcConfigurationProperty",
		reflect.TypeOf((*CfnDeliveryStream_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStreamProps",
		reflect.TypeOf((*CfnDeliveryStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CommonDestinationProps",
		reflect.TypeOf((*CommonDestinationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.CommonDestinationS3Props",
		reflect.TypeOf((*CommonDestinationS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.Compression",
		reflect.TypeOf((*Compression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_Compression{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DataFormatConversionProps",
		reflect.TypeOf((*DataFormatConversionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DataProcessorBindOptions",
		reflect.TypeOf((*DataProcessorBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DataProcessorConfig",
		reflect.TypeOf((*DataProcessorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DataProcessorIdentifier",
		reflect.TypeOf((*DataProcessorIdentifier)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DataProcessorProps",
		reflect.TypeOf((*DataProcessorProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.DeliveryStream",
		reflect.TypeOf((*DeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDeliveryStream)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DeliveryStreamAttributes",
		reflect.TypeOf((*DeliveryStreamAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DeliveryStreamProps",
		reflect.TypeOf((*DeliveryStreamProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DeliveryStreamReference",
		reflect.TypeOf((*DeliveryStreamReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DestinationBindOptions",
		reflect.TypeOf((*DestinationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DestinationConfig",
		reflect.TypeOf((*DestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.DestinationS3BackupProps",
		reflect.TypeOf((*DestinationS3BackupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.DisableLogging",
		reflect.TypeOf((*DisableLogging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
		},
		func() interface{} {
			j := jsiiProxy_DisableLogging{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILoggingConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.EnableLogging",
		reflect.TypeOf((*EnableLogging)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
		},
		func() interface{} {
			j := jsiiProxy_EnableLogging{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ILoggingConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.HiveJsonInputFormat",
		reflect.TypeOf((*HiveJsonInputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createInputFormatConfig", GoMethod: "CreateInputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_HiveJsonInputFormat{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.HiveJsonInputFormatProps",
		reflect.TypeOf((*HiveJsonInputFormatProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.IDataProcessor",
		reflect.TypeOf((*IDataProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			return &jsiiProxy_IDataProcessor{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.IDeliveryStream",
		reflect.TypeOf((*IDeliveryStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamArn", GoGetter: "DeliveryStreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamName", GoGetter: "DeliveryStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutRecords", GoMethod: "GrantPutRecords"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Bytes", GoMethod: "MetricBackupToS3Bytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3DataFreshness", GoMethod: "MetricBackupToS3DataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "metricBackupToS3Records", GoMethod: "MetricBackupToS3Records"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingBytes", GoMethod: "MetricIncomingBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIncomingRecords", GoMethod: "MetricIncomingRecords"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDeliveryStream{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.IDeliveryStreamRef",
		reflect.TypeOf((*IDeliveryStreamRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deliveryStreamRef", GoGetter: "DeliveryStreamRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IDeliveryStreamRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.IDestination",
		reflect.TypeOf((*IDestination)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IDestination{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.IInputFormat",
		reflect.TypeOf((*IInputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createInputFormatConfig", GoMethod: "CreateInputFormatConfig"},
		},
		func() interface{} {
			return &jsiiProxy_IInputFormat{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.ILoggingConfig",
		reflect.TypeOf((*ILoggingConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
		},
		func() interface{} {
			return &jsiiProxy_ILoggingConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.IOutputFormat",
		reflect.TypeOf((*IOutputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createOutputFormatConfig", GoMethod: "CreateOutputFormatConfig"},
		},
		func() interface{} {
			return &jsiiProxy_IOutputFormat{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_kinesisfirehose.ISource",
		reflect.TypeOf((*ISource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
		},
		func() interface{} {
			return &jsiiProxy_ISource{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.InputFormat",
		reflect.TypeOf((*InputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_InputFormat{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.KinesisStreamSource",
		reflect.TypeOf((*KinesisStreamSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisStreamSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.LambdaFunctionProcessor",
		reflect.TypeOf((*LambdaFunctionProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunctionProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDataProcessor)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.OpenXJsonInputFormat",
		reflect.TypeOf((*OpenXJsonInputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createInputFormatConfig", GoMethod: "CreateInputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_OpenXJsonInputFormat{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.OpenXJsonInputFormatProps",
		reflect.TypeOf((*OpenXJsonInputFormatProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.OrcCompression",
		reflect.TypeOf((*OrcCompression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_OrcCompression{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_kinesisfirehose.OrcFormatVersion",
		reflect.TypeOf((*OrcFormatVersion)(nil)).Elem(),
		map[string]interface{}{
			"V0_11": OrcFormatVersion_V0_11,
			"V0_12": OrcFormatVersion_V0_12,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.OrcOutputFormat",
		reflect.TypeOf((*OrcOutputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createOutputFormatConfig", GoMethod: "CreateOutputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_OrcOutputFormat{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOutputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.OrcOutputFormatProps",
		reflect.TypeOf((*OrcOutputFormatProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.OutputFormat",
		reflect.TypeOf((*OutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_OutputFormat{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetCompression",
		reflect.TypeOf((*ParquetCompression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ParquetCompression{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetOutputFormat",
		reflect.TypeOf((*ParquetOutputFormat)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createOutputFormatConfig", GoMethod: "CreateOutputFormatConfig"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
		},
		func() interface{} {
			j := jsiiProxy_ParquetOutputFormat{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOutputFormat)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetOutputFormatProps",
		reflect.TypeOf((*ParquetOutputFormatProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_kinesisfirehose.ParquetWriterVersion",
		reflect.TypeOf((*ParquetWriterVersion)(nil)).Elem(),
		map[string]interface{}{
			"V1": ParquetWriterVersion_V1,
			"V2": ParquetWriterVersion_V2,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.S3Bucket",
		reflect.TypeOf((*S3Bucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Bucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDestination)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.S3BucketProps",
		reflect.TypeOf((*S3BucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.SchemaConfiguration",
		reflect.TypeOf((*SchemaConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_SchemaConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.SchemaConfigurationBindOptions",
		reflect.TypeOf((*SchemaConfigurationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_kinesisfirehose.SchemaConfigurationFromCfnTableProps",
		reflect.TypeOf((*SchemaConfigurationFromCfnTableProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.StreamEncryption",
		reflect.TypeOf((*StreamEncryption)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_StreamEncryption{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_kinesisfirehose.StreamEncryptionType",
		reflect.TypeOf((*StreamEncryptionType)(nil)).Elem(),
		map[string]interface{}{
			"UNENCRYPTED": StreamEncryptionType_UNENCRYPTED,
			"CUSTOMER_MANAGED": StreamEncryptionType_CUSTOMER_MANAGED,
			"AWS_OWNED": StreamEncryptionType_AWS_OWNED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_kinesisfirehose.TimestampParser",
		reflect.TypeOf((*TimestampParser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
		},
		func() interface{} {
			return &jsiiProxy_TimestampParser{}
		},
	)
}
