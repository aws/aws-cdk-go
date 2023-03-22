package awskinesisanalyticsv2


// For an SQL-based Amazon Kinesis Data Analytics application, identifies the Amazon S3 bucket and object that contains the reference data.
//
// A Kinesis Data Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/apiv2/API_UpdateApplication.html) operation to trigger reloading of data into your application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReferenceDataSourceProperty := &s3ReferenceDataSourceProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	fileKey: jsii.String("fileKey"),
//   }
//
type CfnApplicationReferenceDataSource_S3ReferenceDataSourceProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The object key name containing the reference data.
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
}

