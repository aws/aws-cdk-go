package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonOpenSearchServerlessDestinationConfigurationProperty := &AmazonOpenSearchServerlessDestinationConfigurationProperty{
//   	IndexName: jsii.String("indexName"),
//   	RoleArn: jsii.String("roleArn"),
//   	S3Configuration: &S3DestinationConfigurationProperty{
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
//
//   	// the properties below are optional
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
//   	ProcessingConfiguration: &ProcessingConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Processors: []interface{}{
//   			&ProcessorProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				Parameters: []interface{}{
//   					&ProcessorParameterProperty{
//   						ParameterName: jsii.String("parameterName"),
//   						ParameterValue: jsii.String("parameterValue"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	RetryOptions: &AmazonOpenSearchServerlessRetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   	S3BackupMode: jsii.String("s3BackupMode"),
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
type CfnDeliveryStream_AmazonOpenSearchServerlessDestinationConfigurationProperty struct {
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.IndexName`.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.RoleARN`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.BufferingHints`.
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.CloudWatchLoggingOptions`.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.CollectionEndpoint`.
	CollectionEndpoint *string `field:"optional" json:"collectionEndpoint" yaml:"collectionEndpoint"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.ProcessingConfiguration`.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.RetryOptions`.
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.S3BackupMode`.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// `CfnDeliveryStream.AmazonOpenSearchServerlessDestinationConfigurationProperty.VpcConfiguration`.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

