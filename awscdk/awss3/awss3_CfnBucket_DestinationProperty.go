package awss3


// Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &destinationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	format: jsii.String("format"),
//
//   	// the properties below are optional
//   	bucketAccountId: jsii.String("bucketAccountId"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnBucket_DestinationProperty struct {
	// The Amazon Resource Name (ARN) of the bucket to which data is exported.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Specifies the file format used when exporting data to Amazon S3.
	//
	// *Allowed values* : `CSV` | `ORC` | `Parquet`.
	Format *string `field:"required" json:"format" yaml:"format"`
	// The account ID that owns the destination S3 bucket.
	//
	// If no account ID is provided, the owner is not validated before exporting data.
	//
	// > Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The prefix to use when exporting data.
	//
	// The prefix is prepended to all results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

