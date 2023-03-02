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
//   dataExportProperty := &dataExportProperty{
//   	cloudWatchMetrics: &cloudWatchMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	s3BucketDestination: &s3BucketDestinationProperty{
//   		accountId: jsii.String("accountId"),
//   		arn: jsii.String("arn"),
//   		format: jsii.String("format"),
//   		outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   		// the properties below are optional
//   		encryption: &encryptionProperty{
//   			ssekms: &sSEKMSProperty{
//   				keyId: jsii.String("keyId"),
//   			},
//   			sses3: sses3,
//   		},
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnStorageLens_DataExportProperty struct {
	// This property enables the Amazon CloudWatch publishing option for S3 Storage Lens metrics.
	CloudWatchMetrics interface{} `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// This property contains the details of the bucket where the S3 Storage Lens metrics export will be placed.
	S3BucketDestination interface{} `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
}

