package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeDestinationConfigurationProperty := &SnowflakeDestinationConfigurationProperty{
//   	AccountUrl: jsii.String("accountUrl"),
//   	Database: jsii.String("database"),
//   	PrivateKey: jsii.String("privateKey"),
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
//   	Schema: jsii.String("schema"),
//   	Table: jsii.String("table"),
//   	User: jsii.String("user"),
//
//   	// the properties below are optional
//   	CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LogGroupName: jsii.String("logGroupName"),
//   		LogStreamName: jsii.String("logStreamName"),
//   	},
//   	ContentColumnName: jsii.String("contentColumnName"),
//   	DataLoadingOption: jsii.String("dataLoadingOption"),
//   	KeyPassphrase: jsii.String("keyPassphrase"),
//   	MetaDataColumnName: jsii.String("metaDataColumnName"),
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
//   	RetryOptions: &SnowflakeRetryOptionsProperty{
//   		DurationInSeconds: jsii.Number(123),
//   	},
//   	S3BackupMode: jsii.String("s3BackupMode"),
//   	SnowflakeRoleConfiguration: &SnowflakeRoleConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		SnowflakeRole: jsii.String("snowflakeRole"),
//   	},
//   	SnowflakeVpcConfiguration: &SnowflakeVpcConfigurationProperty{
//   		PrivateLinkVpceId: jsii.String("privateLinkVpceId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html
//
type CfnDeliveryStream_SnowflakeDestinationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-accounturl
	//
	AccountUrl *string `field:"required" json:"accountUrl" yaml:"accountUrl"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-privatekey
	//
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-s3configuration
	//
	S3Configuration interface{} `field:"required" json:"s3Configuration" yaml:"s3Configuration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-schema
	//
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-table
	//
	Table *string `field:"required" json:"table" yaml:"table"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-user
	//
	User *string `field:"required" json:"user" yaml:"user"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-cloudwatchloggingoptions
	//
	CloudWatchLoggingOptions interface{} `field:"optional" json:"cloudWatchLoggingOptions" yaml:"cloudWatchLoggingOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-contentcolumnname
	//
	ContentColumnName *string `field:"optional" json:"contentColumnName" yaml:"contentColumnName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-dataloadingoption
	//
	DataLoadingOption *string `field:"optional" json:"dataLoadingOption" yaml:"dataLoadingOption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-keypassphrase
	//
	KeyPassphrase *string `field:"optional" json:"keyPassphrase" yaml:"keyPassphrase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-metadatacolumnname
	//
	MetaDataColumnName *string `field:"optional" json:"metaDataColumnName" yaml:"metaDataColumnName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-processingconfiguration
	//
	ProcessingConfiguration interface{} `field:"optional" json:"processingConfiguration" yaml:"processingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-retryoptions
	//
	RetryOptions interface{} `field:"optional" json:"retryOptions" yaml:"retryOptions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-s3backupmode
	//
	S3BackupMode *string `field:"optional" json:"s3BackupMode" yaml:"s3BackupMode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-snowflakeroleconfiguration
	//
	SnowflakeRoleConfiguration interface{} `field:"optional" json:"snowflakeRoleConfiguration" yaml:"snowflakeRoleConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-snowflakedestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-snowflakedestinationconfiguration-snowflakevpcconfiguration
	//
	SnowflakeVpcConfiguration interface{} `field:"optional" json:"snowflakeVpcConfiguration" yaml:"snowflakeVpcConfiguration"`
}

