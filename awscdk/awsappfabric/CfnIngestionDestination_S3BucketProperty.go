package awsappfabric


// Contains information about an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketProperty := &S3BucketProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-s3bucket.html
//
type CfnIngestionDestination_S3BucketProperty struct {
	// The name of the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-s3bucket.html#cfn-appfabric-ingestiondestination-s3bucket-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appfabric-ingestiondestination-s3bucket.html#cfn-appfabric-ingestiondestination-s3bucket-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

