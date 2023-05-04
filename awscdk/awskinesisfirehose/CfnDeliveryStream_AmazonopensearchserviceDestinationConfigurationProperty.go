package awskinesisfirehose


// Describes the configuration of a destination in Amazon OpenSearch Service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonopensearchserviceDestinationConfigurationProperty := &AmazonopensearchserviceDestinationConfigurationProperty{
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
//   	BufferingHints: &AmazonopensearchserviceBufferingHintsProperty{
//   		IntervalInSeconds: jsii.Number(123),
//   		SizeInMBs: jsii.Number(123),
//   	},
//   	CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamName: jsii.String("logStreamName"),
//   	},
//   	ClusterEndpoint: jsii.String("clusterEndpoint"),
//   	DomainArn: jsii.String("domainArn"),
//   	IndexRotationPeriod: jsii.String("indexRotationPeriod"),
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
//   	RetryOptions: &AmazonopensearchserviceRetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   	S3BackupMode: jsii.String("s3BackupMode"),
//   	TypeName: jsii.String("typeName"),
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
type CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty struct {
	// The Amazon OpenSearch Service index name.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The Amazon Resource Name (ARN) of the IAM role to be assumed by Kinesis Data Firehose for calling the Amazon OpenSearch Service Configuration API and for indexing documents.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Describes the configuration of a destination in Amazon S3.
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// The buffering options.
	//
	// If no value is specified, the default values for AmazonopensearchserviceBufferingHints are used.
	BufferingHints interface{} `field:"optional" json:"bufferingHints" yaml:"bufferingHints"`
	// Describes the Amazon CloudWatch logging options for your delivery stream.
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// The endpoint to use when communicating with the cluster.
	//
	// Specify either this ClusterEndpoint or the DomainARN field.
	ClusterEndpoint *string `field:"optional" json:"clusterEndpoint" yaml:"clusterEndpoint"`
	// The ARN of the Amazon OpenSearch Service domain.
	DomainArn *string `field:"optional" json:"domainArn" yaml:"domainArn"`
	// The Amazon OpenSearch Service index rotation period.
	//
	// Index rotation appends a timestamp to the IndexName to facilitate the expiration of old data.
	IndexRotationPeriod *string `field:"optional" json:"indexRotationPeriod" yaml:"indexRotationPeriod"`
	// Describes a data processing configuration.
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// The retry behavior in case Kinesis Data Firehose is unable to deliver documents to Amazon OpenSearch Service.
	//
	// The default value is 300 (5 minutes).
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// Defines how documents should be delivered to Amazon S3.
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// The Amazon OpenSearch Service type name.
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The details of the VPC of the Amazon OpenSearch Service destination.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

