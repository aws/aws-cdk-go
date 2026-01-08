package awss3


// This resource contains the details of the Amazon S3 Storage Lens metrics export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   dataExportProperty := &DataExportProperty{
//   	CloudWatchMetrics: &CloudWatchMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	S3BucketDestination: &S3BucketDestinationProperty{
//   		AccountId: jsii.String("accountId"),
//   		Arn: jsii.String("arn"),
//   		Format: jsii.String("format"),
//   		OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   		// the properties below are optional
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   		Prefix: jsii.String("prefix"),
//   	},
//   	StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   		IsEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html
//
type CfnStorageLens_DataExportProperty struct {
	// This property enables the Amazon CloudWatch publishing option for S3 Storage Lens metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html#cfn-s3-storagelens-dataexport-cloudwatchmetrics
	//
	CloudWatchMetrics interface{} `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// This property contains the details of the bucket where the S3 Storage Lens metrics export will be placed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html#cfn-s3-storagelens-dataexport-s3bucketdestination
	//
	S3BucketDestination interface{} `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// This property contains the details of the S3 table bucket where the S3 Storage Lens default metrics report will be placed.
	//
	// This property enables you to store your Storage Lens metrics in read-only S3 Tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html#cfn-s3-storagelens-dataexport-storagelenstabledestination
	//
	StorageLensTableDestination interface{} `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

