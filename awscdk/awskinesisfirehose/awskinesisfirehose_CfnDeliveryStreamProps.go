package awskinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDeliveryStream`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryStreamProps := &CfnDeliveryStreamProps{
//   	AmazonOpenSearchServerlessDestinationConfiguration: &AmazonOpenSearchServerlessDestinationConfigurationProperty{
//   		IndexName: jsii.String("indexName"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &AmazonOpenSearchServerlessBufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CollectionEndpoint: jsii.String("collectionEndpoint"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &AmazonOpenSearchServerlessRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	AmazonopensearchserviceDestinationConfiguration: &AmazonopensearchserviceDestinationConfigurationProperty{
//   		IndexName: jsii.String("indexName"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &AmazonopensearchserviceBufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ClusterEndpoint: jsii.String("clusterEndpoint"),
//   		DocumentIdOptions: &DocumentIdOptionsProperty{
//   			DefaultDocumentIdFormat: jsii.String("defaultDocumentIdFormat"),
//   		},
//   		DomainArn: jsii.String("domainArn"),
//   		IndexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &AmazonopensearchserviceRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   		TypeName: jsii.String("typeName"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	DeliveryStreamEncryptionConfigurationInput: &DeliveryStreamEncryptionConfigurationInputProperty{
//   		KeyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		KeyArn: jsii.String("keyArn"),
//   	},
//   	DeliveryStreamName: jsii.String("deliveryStreamName"),
//   	DeliveryStreamType: jsii.String("deliveryStreamType"),
//   	ElasticsearchDestinationConfiguration: &ElasticsearchDestinationConfigurationProperty{
//   		IndexName: jsii.String("indexName"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &ElasticsearchBufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ClusterEndpoint: jsii.String("clusterEndpoint"),
//   		DocumentIdOptions: &DocumentIdOptionsProperty{
//   			DefaultDocumentIdFormat: jsii.String("defaultDocumentIdFormat"),
//   		},
//   		DomainArn: jsii.String("domainArn"),
//   		IndexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &ElasticsearchRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   		TypeName: jsii.String("typeName"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	ExtendedS3DestinationConfiguration: &ExtendedS3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CompressionFormat: jsii.String("compressionFormat"),
//   		DataFormatConversionConfiguration: &DataFormatConversionConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			InputFormatConfiguration: &InputFormatConfigurationProperty{
//   				Deserializer: &DeserializerProperty{
//   					HiveJsonSerDe: &HiveJsonSerDeProperty{
//   						TimestampFormats: []*string{
//   							jsii.String("timestampFormats"),
//   						},
//   					},
//   					OpenXJsonSerDe: &OpenXJsonSerDeProperty{
//   						CaseInsensitive: jsii.Boolean(false),
//   						ColumnToJsonKeyMappings: map[string]*string{
//   							"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   						},
//   						ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OutputFormatConfiguration: &OutputFormatConfigurationProperty{
//   				Serializer: &SerializerProperty{
//   					OrcSerDe: &OrcSerDeProperty{
//   						BlockSizeBytes: jsii.Number(123),
//   						BloomFilterColumns: []*string{
//   							jsii.String("bloomFilterColumns"),
//   						},
//   						BloomFilterFalsePositiveProbability: jsii.Number(123),
//   						Compression: jsii.String("compression"),
//   						DictionaryKeyThreshold: jsii.Number(123),
//   						EnablePadding: jsii.Boolean(false),
//   						FormatVersion: jsii.String("formatVersion"),
//   						PaddingTolerance: jsii.Number(123),
//   						RowIndexStride: jsii.Number(123),
//   						StripeSizeBytes: jsii.Number(123),
//   					},
//   					ParquetSerDe: &ParquetSerDeProperty{
//   						BlockSizeBytes: jsii.Number(123),
//   						Compression: jsii.String("compression"),
//   						EnableDictionaryCompression: jsii.Boolean(false),
//   						MaxPaddingBytes: jsii.Number(123),
//   						PageSizeBytes: jsii.Number(123),
//   						WriterVersion: jsii.String("writerVersion"),
//   					},
//   				},
//   			},
//   			SchemaConfiguration: &SchemaConfigurationProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				DatabaseName: jsii.String("databaseName"),
//   				Region: jsii.String("region"),
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//   				VersionId: jsii.String("versionId"),
//   			},
//   		},
//   		DynamicPartitioningConfiguration: &DynamicPartitioningConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			RetryOptions: &RetryOptionsProperty{
//   				DurationInSeconds: jsii.Number(123),
//   			},
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   				AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Prefix: jsii.String("prefix"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		S3BackupConfiguration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	HttpEndpointDestinationConfiguration: &HttpEndpointDestinationConfigurationProperty{
//   		EndpointConfiguration: &HttpEndpointConfigurationProperty{
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			AccessKey: jsii.String("accessKey"),
//   			Name: jsii.String("name"),
//   		},
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RequestConfiguration: &HttpEndpointRequestConfigurationProperty{
//   			CommonAttributes: []interface{}{
//   				&HttpEndpointCommonAttributeProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					AttributeValue: jsii.String("attributeValue"),
//   				},
//   			},
//   			ContentEncoding: jsii.String("contentEncoding"),
//   		},
//   		RetryOptions: &RetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	KinesisStreamSourceConfiguration: &KinesisStreamSourceConfigurationProperty{
//   		KinesisStreamArn: jsii.String("kinesisStreamArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	RedshiftDestinationConfiguration: &RedshiftDestinationConfigurationProperty{
//   		ClusterJdbcurl: jsii.String("clusterJdbcurl"),
//   		CopyCommand: &CopyCommandProperty{
//   			DataTableName: jsii.String("dataTableName"),
//
//   			// the properties below are optional
//   			CopyOptions: jsii.String("copyOptions"),
//   			DataTableColumns: jsii.String("dataTableColumns"),
//   		},
//   		Password: jsii.String("password"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		Username: jsii.String("username"),
//
//   		// the properties below are optional
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &RedshiftRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupConfiguration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CompressionFormat: jsii.String("compressionFormat"),
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   				AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	SplunkDestinationConfiguration: &SplunkDestinationConfigurationProperty{
//   		HecEndpoint: jsii.String("hecEndpoint"),
//   		HecEndpointType: jsii.String("hecEndpointType"),
//   		HecToken: jsii.String("hecToken"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		HecAcknowledgmentTimeoutInSeconds: jsii.Number(123),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &SplunkRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeliveryStreamProps struct {
	// `AWS::KinesisFirehose::DeliveryStream.AmazonOpenSearchServerlessDestinationConfiguration`.
	AmazonOpenSearchServerlessDestinationConfiguration interface{} `field:"optional" json:"amazonOpenSearchServerlessDestinationConfiguration" yaml:"amazonOpenSearchServerlessDestinationConfiguration"`
	// The destination in Amazon OpenSearch Service.
	//
	// You can specify only one destination.
	AmazonopensearchserviceDestinationConfiguration interface{} `field:"optional" json:"amazonopensearchserviceDestinationConfiguration" yaml:"amazonopensearchserviceDestinationConfiguration"`
	// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput interface{} `field:"optional" json:"deliveryStreamEncryptionConfigurationInput" yaml:"deliveryStreamEncryptionConfigurationInput"`
	// The name of the delivery stream.
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// The delivery stream type. This can be one of the following values:.
	//
	// - `DirectPut` : Provider applications access the delivery stream directly.
	// - `KinesisStreamAsSource` : The delivery stream uses a Kinesis data stream as a source.
	DeliveryStreamType *string `field:"optional" json:"deliveryStreamType" yaml:"deliveryStreamType"`
	// An Amazon ES destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon ES destination to an Amazon S3 or Amazon Redshift destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ElasticsearchDestinationConfiguration interface{} `field:"optional" json:"elasticsearchDestinationConfiguration" yaml:"elasticsearchDestinationConfiguration"`
	// An Amazon S3 destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Extended S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ExtendedS3DestinationConfiguration interface{} `field:"optional" json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.
	//
	// You can specify only one destination.
	HttpEndpointDestinationConfiguration interface{} `field:"optional" json:"httpEndpointDestinationConfiguration" yaml:"httpEndpointDestinationConfiguration"`
	// When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html) containing the Kinesis stream ARN and the role ARN for the source stream.
	KinesisStreamSourceConfiguration interface{} `field:"optional" json:"kinesisStreamSourceConfiguration" yaml:"kinesisStreamSourceConfiguration"`
	// An Amazon Redshift destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Redshift destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	RedshiftDestinationConfiguration interface{} `field:"optional" json:"redshiftDestinationConfiguration" yaml:"redshiftDestinationConfiguration"`
	// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	S3DestinationConfiguration interface{} `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	// The configuration of a destination in Splunk for the delivery stream.
	SplunkDestinationConfiguration interface{} `field:"optional" json:"splunkDestinationConfiguration" yaml:"splunkDestinationConfiguration"`
	// A set of tags to assign to the delivery stream.
	//
	// A tag is a key-value pair that you can define and assign to AWS resources. Tags are metadata. For example, you can add friendly names and descriptions or other types of information that can help you distinguish the delivery stream. For more information about tags, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

