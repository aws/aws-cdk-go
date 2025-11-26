package previewawskinesisfirehosemixins


// Describes the configuration of a destination in the Serverless offering for Amazon OpenSearch Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   amazonOpenSearchServerlessDestinationConfigurationProperty := &AmazonOpenSearchServerlessDestinationConfigurationProperty{
//   	BufferingHints: &AmazonOpenSearchServerlessBufferingHintsProperty{
//   		IntervalInSeconds: jsii.Number(123),
//   		SizeInMBs: jsii.Number(123),
//   	},
//   	CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamName: jsii.String("logStreamName"),
//   	},
//   	CollectionEndpoint: jsii.String("collectionEndpoint"),
//   	IndexName: jsii.String("indexName"),
//   	ProcessingConfiguration: &ProcessingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Processors: []interface{}{
//   			&ProcessorProperty{
//   				Parameters: []interface{}{
//   					&ProcessorParameterProperty{
//   						ParameterName: jsii.String("parameterName"),
//   						ParameterValue: jsii.String("parameterValue"),
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	RetryOptions: &AmazonOpenSearchServerlessRetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	S3BackupMode: jsii.String("s3BackupMode"),
//   	S3Configuration: &S3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
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
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html
//
type CfnDeliveryStreamPropsMixin_AmazonOpenSearchServerlessDestinationConfigurationProperty struct {
	// The buffering options.
	//
	// If no value is specified, the default values for AmazonopensearchserviceBufferingHints are used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-bufferinghints
	//
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-cloudwatchloggingoptions
	//
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The endpoint to use when communicating with the collection in the Serverless offering for Amazon OpenSearch Service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-collectionendpoint
	//
	CollectionEndpoint *string `field:"optional" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// The Serverless offering for Amazon OpenSearch Service index name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-processingconfiguration
	//
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior in case Firehose is unable to deliver documents to the Serverless offering for Amazon OpenSearch Service.
	//
	// The default value is 300 (5 minutes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-retryoptions
	//
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// The Amazon Resource Name (ARN) of the IAM role to be assumed by Firehose for calling the Serverless offering for Amazon OpenSearch Service Configuration API and for indexing documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Defines how documents should be delivered to Amazon S3.
	//
	// When it is set to FailedDocumentsOnly, Firehose writes any documents that could not be indexed to the configured Amazon S3 destination, with AmazonOpenSearchService-failed/ appended to the key prefix. When set to AllDocuments, Firehose delivers all incoming records to Amazon S3, and also writes failed documents with AmazonOpenSearchService-failed/ appended to the prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-s3backupmode
	//
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-amazonopensearchserverlessdestinationconfiguration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

