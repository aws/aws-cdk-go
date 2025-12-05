package previewawss3mixins


// This resource contains the details of the Amazon S3 Storage Lens metrics export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   		Format: jsii.String("format"),
//   		OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	StorageLensTableDestination: &StorageLensTableDestinationProperty{
//   		Encryption: &EncryptionProperty{
//   			Ssekms: &SSEKMSProperty{
//   				KeyId: jsii.String("keyId"),
//   			},
//   			Sses3: sses3,
//   		},
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html
//
type CfnStorageLensPropsMixin_DataExportProperty struct {
	// This property enables the Amazon CloudWatch publishing option for S3 Storage Lens metrics.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html#cfn-s3-storagelens-dataexport-cloudwatchmetrics
	//
	CloudWatchMetrics interface{} `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// This property contains the details of the bucket where the S3 Storage Lens metrics export will be placed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html#cfn-s3-storagelens-dataexport-s3bucketdestination
	//
	S3BucketDestination interface{} `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
	// S3 Tables destination settings for the Amazon S3 Storage Lens metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelens-dataexport.html#cfn-s3-storagelens-dataexport-storagelenstabledestination
	//
	StorageLensTableDestination interface{} `field:"optional" json:"storageLensTableDestination" yaml:"storageLensTableDestination"`
}

