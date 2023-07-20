package awskinesisanalytics


// Identifies the S3 bucket and object that contains the reference data.
//
// Also identifies the IAM role Amazon Kinesis Analytics can assume to read this object on your behalf.
//
// An Amazon Kinesis Analytics application loads reference data only once. If the data changes, you call the [UpdateApplication](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_UpdateApplication.html) operation to trigger reloading of data into your application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ReferenceDataSourceProperty := &S3ReferenceDataSourceProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	FileKey: jsii.String("fileKey"),
//   	ReferenceRoleArn: jsii.String("referenceRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html
//
type CfnApplicationReferenceDataSource_S3ReferenceDataSourceProperty struct {
	// Amazon Resource Name (ARN) of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-s3referencedatasource-bucketarn
	//
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Object key name containing reference data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-s3referencedatasource-filekey
	//
	FileKey *string `field:"required" json:"fileKey" yaml:"fileKey"`
	// ARN of the IAM role that the service can assume to read data on your behalf.
	//
	// This role must have permission for the `s3:GetObject` action on the object and trust policy that allows Amazon Kinesis Analytics service principal to assume this role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-s3referencedatasource-referencerolearn
	//
	ReferenceRoleArn *string `field:"required" json:"referenceRoleArn" yaml:"referenceRoleArn"`
}

