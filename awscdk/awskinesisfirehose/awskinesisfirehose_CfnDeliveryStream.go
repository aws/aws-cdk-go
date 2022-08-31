package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::KinesisFirehose::DeliveryStream`.
//
// The `AWS::KinesisFirehose::DeliveryStream` resource specifies an Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivery stream that delivers real-time streaming data to an Amazon Simple Storage Service (Amazon S3), Amazon Redshift, or Amazon Elasticsearch Service (Amazon ES) destination. For more information, see [Creating an Amazon Kinesis Data Firehose Delivery Stream](https://docs.aws.amazon.com/firehose/latest/dev/basic-create.html) in the *Amazon Kinesis Data Firehose Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryStream := awscdk.Aws_kinesisfirehose.NewCfnDeliveryStream(this, jsii.String("MyCfnDeliveryStream"), &cfnDeliveryStreamProps{
//   	amazonopensearchserviceDestinationConfiguration: &amazonopensearchserviceDestinationConfigurationProperty{
//   		indexName: jsii.String("indexName"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		bufferingHints: &amazonopensearchserviceBufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		clusterEndpoint: jsii.String("clusterEndpoint"),
//   		domainArn: jsii.String("domainArn"),
//   		indexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &amazonopensearchserviceRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   		typeName: jsii.String("typeName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	deliveryStreamEncryptionConfigurationInput: &deliveryStreamEncryptionConfigurationInputProperty{
//   		keyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		keyArn: jsii.String("keyArn"),
//   	},
//   	deliveryStreamName: jsii.String("deliveryStreamName"),
//   	deliveryStreamType: jsii.String("deliveryStreamType"),
//   	elasticsearchDestinationConfiguration: &elasticsearchDestinationConfigurationProperty{
//   		indexName: jsii.String("indexName"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		bufferingHints: &elasticsearchBufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		clusterEndpoint: jsii.String("clusterEndpoint"),
//   		domainArn: jsii.String("domainArn"),
//   		indexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &elasticsearchRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   		typeName: jsii.String("typeName"),
//   		vpcConfiguration: &vpcConfigurationProperty{
//   			roleArn: jsii.String("roleArn"),
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	extendedS3DestinationConfiguration: &extendedS3DestinationConfigurationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		bufferingHints: &bufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		compressionFormat: jsii.String("compressionFormat"),
//   		dataFormatConversionConfiguration: &dataFormatConversionConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			inputFormatConfiguration: &inputFormatConfigurationProperty{
//   				deserializer: &deserializerProperty{
//   					hiveJsonSerDe: &hiveJsonSerDeProperty{
//   						timestampFormats: []*string{
//   							jsii.String("timestampFormats"),
//   						},
//   					},
//   					openXJsonSerDe: &openXJsonSerDeProperty{
//   						caseInsensitive: jsii.Boolean(false),
//   						columnToJsonKeyMappings: map[string]*string{
//   							"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   						},
//   						convertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			outputFormatConfiguration: &outputFormatConfigurationProperty{
//   				serializer: &serializerProperty{
//   					orcSerDe: &orcSerDeProperty{
//   						blockSizeBytes: jsii.Number(123),
//   						bloomFilterColumns: []*string{
//   							jsii.String("bloomFilterColumns"),
//   						},
//   						bloomFilterFalsePositiveProbability: jsii.Number(123),
//   						compression: jsii.String("compression"),
//   						dictionaryKeyThreshold: jsii.Number(123),
//   						enablePadding: jsii.Boolean(false),
//   						formatVersion: jsii.String("formatVersion"),
//   						paddingTolerance: jsii.Number(123),
//   						rowIndexStride: jsii.Number(123),
//   						stripeSizeBytes: jsii.Number(123),
//   					},
//   					parquetSerDe: &parquetSerDeProperty{
//   						blockSizeBytes: jsii.Number(123),
//   						compression: jsii.String("compression"),
//   						enableDictionaryCompression: jsii.Boolean(false),
//   						maxPaddingBytes: jsii.Number(123),
//   						pageSizeBytes: jsii.Number(123),
//   						writerVersion: jsii.String("writerVersion"),
//   					},
//   				},
//   			},
//   			schemaConfiguration: &schemaConfigurationProperty{
//   				catalogId: jsii.String("catalogId"),
//   				databaseName: jsii.String("databaseName"),
//   				region: jsii.String("region"),
//   				roleArn: jsii.String("roleArn"),
//   				tableName: jsii.String("tableName"),
//   				versionId: jsii.String("versionId"),
//   			},
//   		},
//   		dynamicPartitioningConfiguration: &dynamicPartitioningConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			retryOptions: &retryOptionsProperty{
//   				durationInSeconds: jsii.Number(123),
//   			},
//   		},
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   				awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		prefix: jsii.String("prefix"),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		s3BackupConfiguration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	httpEndpointDestinationConfiguration: &httpEndpointDestinationConfigurationProperty{
//   		endpointConfiguration: &httpEndpointConfigurationProperty{
//   			url: jsii.String("url"),
//
//   			// the properties below are optional
//   			accessKey: jsii.String("accessKey"),
//   			name: jsii.String("name"),
//   		},
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		bufferingHints: &bufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		requestConfiguration: &httpEndpointRequestConfigurationProperty{
//   			commonAttributes: []interface{}{
//   				&httpEndpointCommonAttributeProperty{
//   					attributeName: jsii.String("attributeName"),
//   					attributeValue: jsii.String("attributeValue"),
//   				},
//   			},
//   			contentEncoding: jsii.String("contentEncoding"),
//   		},
//   		retryOptions: &retryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		roleArn: jsii.String("roleArn"),
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	kinesisStreamSourceConfiguration: &kinesisStreamSourceConfigurationProperty{
//   		kinesisStreamArn: jsii.String("kinesisStreamArn"),
//   		roleArn: jsii.String("roleArn"),
//   	},
//   	redshiftDestinationConfiguration: &redshiftDestinationConfigurationProperty{
//   		clusterJdbcurl: jsii.String("clusterJdbcurl"),
//   		copyCommand: &copyCommandProperty{
//   			dataTableName: jsii.String("dataTableName"),
//
//   			// the properties below are optional
//   			copyOptions: jsii.String("copyOptions"),
//   			dataTableColumns: jsii.String("dataTableColumns"),
//   		},
//   		password: jsii.String("password"),
//   		roleArn: jsii.String("roleArn"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		username: jsii.String("username"),
//
//   		// the properties below are optional
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &redshiftRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupConfiguration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	s3DestinationConfiguration: &s3DestinationConfigurationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		bufferingHints: &bufferingHintsProperty{
//   			intervalInSeconds: jsii.Number(123),
//   			sizeInMBs: jsii.Number(123),
//   		},
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		compressionFormat: jsii.String("compressionFormat"),
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   				awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		prefix: jsii.String("prefix"),
//   	},
//   	splunkDestinationConfiguration: &splunkDestinationConfigurationProperty{
//   		hecEndpoint: jsii.String("hecEndpoint"),
//   		hecEndpointType: jsii.String("hecEndpointType"),
//   		hecToken: jsii.String("hecToken"),
//   		s3Configuration: &s3DestinationConfigurationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			bufferingHints: &bufferingHintsProperty{
//   				intervalInSeconds: jsii.Number(123),
//   				sizeInMBs: jsii.Number(123),
//   			},
//   			cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   				enabled: jsii.Boolean(false),
//   				logGroupName: jsii.String("logGroupName"),
//   				logStreamName: jsii.String("logStreamName"),
//   			},
//   			compressionFormat: jsii.String("compressionFormat"),
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				kmsEncryptionConfig: &kMSEncryptionConfigProperty{
//   					awskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				noEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		cloudWatchLoggingOptions: &cloudWatchLoggingOptionsProperty{
//   			enabled: jsii.Boolean(false),
//   			logGroupName: jsii.String("logGroupName"),
//   			logStreamName: jsii.String("logStreamName"),
//   		},
//   		hecAcknowledgmentTimeoutInSeconds: jsii.Number(123),
//   		processingConfiguration: &processingConfigurationProperty{
//   			enabled: jsii.Boolean(false),
//   			processors: []interface{}{
//   				&processorProperty{
//   					type: jsii.String("type"),
//
//   					// the properties below are optional
//   					parameters: []interface{}{
//   						&processorParameterProperty{
//   							parameterName: jsii.String("parameterName"),
//   							parameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		retryOptions: &splunkRetryOptionsProperty{
//   			durationInSeconds: jsii.Number(123),
//   		},
//   		s3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDeliveryStream interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The destination in Amazon OpenSearch Service.
	//
	// You can specify only one destination.
	AmazonopensearchserviceDestinationConfiguration() interface{}
	SetAmazonopensearchserviceDestinationConfiguration(val interface{})
	// The Amazon Resource Name (ARN) of the delivery stream, such as `arn:aws:firehose:us-east-2:123456789012:deliverystream/delivery-stream-name` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput() interface{}
	SetDeliveryStreamEncryptionConfigurationInput(val interface{})
	// The name of the delivery stream.
	DeliveryStreamName() *string
	SetDeliveryStreamName(val *string)
	// The delivery stream type. This can be one of the following values:.
	//
	// - `DirectPut` : Provider applications access the delivery stream directly.
	// - `KinesisStreamAsSource` : The delivery stream uses a Kinesis data stream as a source.
	DeliveryStreamType() *string
	SetDeliveryStreamType(val *string)
	// An Amazon ES destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon ES destination to an Amazon S3 or Amazon Redshift destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ElasticsearchDestinationConfiguration() interface{}
	SetElasticsearchDestinationConfiguration(val interface{})
	// An Amazon S3 destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Extended S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ExtendedS3DestinationConfiguration() interface{}
	SetExtendedS3DestinationConfiguration(val interface{})
	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.
	//
	// You can specify only one destination.
	HttpEndpointDestinationConfiguration() interface{}
	SetHttpEndpointDestinationConfiguration(val interface{})
	// When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html) containing the Kinesis stream ARN and the role ARN for the source stream.
	KinesisStreamSourceConfiguration() interface{}
	SetKinesisStreamSourceConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// An Amazon Redshift destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Redshift destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	RedshiftDestinationConfiguration() interface{}
	SetRedshiftDestinationConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	S3DestinationConfiguration() interface{}
	SetS3DestinationConfiguration(val interface{})
	// The configuration of a destination in Splunk for the delivery stream.
	SplunkDestinationConfiguration() interface{}
	SetSplunkDestinationConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A set of tags to assign to the delivery stream.
	//
	// A tag is a key-value pair that you can define and assign to AWS resources. Tags are metadata. For example, you can add friendly names and descriptions or other types of information that can help you distinguish the delivery stream. For more information about tags, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDeliveryStream
type jsiiProxy_CfnDeliveryStream struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeliveryStream) AmazonopensearchserviceDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) ElasticsearchDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) ExtendedS3DestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendedS3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) HttpEndpointDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpEndpointDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) KinesisStreamSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) RedshiftDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) S3DestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) SplunkDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisFirehose::DeliveryStream`.
func NewCfnDeliveryStream(scope awscdk.Construct, id *string, props *CfnDeliveryStreamProps) CfnDeliveryStream {
	_init_.Initialize()

	j := jsiiProxy_CfnDeliveryStream{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisFirehose::DeliveryStream`.
func NewCfnDeliveryStream_Override(c CfnDeliveryStream, scope awscdk.Construct, id *string, props *CfnDeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetAmazonopensearchserviceDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetDeliveryStreamEncryptionConfigurationInput(val interface{}) {
	_jsii_.Set(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetDeliveryStreamName(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamName",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetDeliveryStreamType(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamType",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetElasticsearchDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetExtendedS3DestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"extendedS3DestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetHttpEndpointDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"httpEndpointDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetKinesisStreamSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"kinesisStreamSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetRedshiftDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"redshiftDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetS3DestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"s3DestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream) SetSplunkDestinationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"splunkDestinationConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDeliveryStream_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeliveryStream_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeliveryStream_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_kinesisfirehose.CfnDeliveryStream",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeliveryStream) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDeliveryStream) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

