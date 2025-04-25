package awskinesisanalyticsv2


// For an SQL-based Amazon Kinesis Data Analytics application, identifies the Amazon S3 bucket and object that contains the reference data.
//
// A Kinesis Data Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/managed-flink/latest/apiv2/API_UpdateApplication.html) operation to trigger reloading of data into your application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReferenceDataSourceProperty := &S3ReferenceDataSourceProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	FileKey: jsii.String("fileKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-s3referencedatasource.html
//
type CfnApplicationReferenceDataSource_S3ReferenceDataSourceProperty struct {
	// The Amazon Resource Name (ARN) of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-s3referencedatasource-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// The object key name containing the reference data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalyticsv2-applicationreferencedatasource-s3referencedatasource-filekey
	//
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
}

