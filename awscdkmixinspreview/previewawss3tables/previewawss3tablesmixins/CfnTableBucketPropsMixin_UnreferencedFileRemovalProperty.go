package previewawss3tablesmixins


// The unreferenced file removal settings for your table bucket.
//
// Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. For more information, see the [*Amazon S3 User Guide*](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-table-buckets-maintenance.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   unreferencedFileRemovalProperty := &UnreferencedFileRemovalProperty{
//   	NoncurrentDays: jsii.Number(123),
//   	Status: jsii.String("status"),
//   	UnreferencedDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-unreferencedfileremoval.html
//
type CfnTableBucketPropsMixin_UnreferencedFileRemovalProperty struct {
	// The number of days an object can be noncurrent before Amazon S3 deletes it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-unreferencedfileremoval.html#cfn-s3tables-tablebucket-unreferencedfileremoval-noncurrentdays
	//
	NoncurrentDays *float64 `field:"optional" json:"noncurrentDays" yaml:"noncurrentDays"`
	// The status of the unreferenced file removal configuration for your table bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-unreferencedfileremoval.html#cfn-s3tables-tablebucket-unreferencedfileremoval-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The number of days an object must be unreferenced by your table before Amazon S3 marks the object as noncurrent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3tables-tablebucket-unreferencedfileremoval.html#cfn-s3tables-tablebucket-unreferencedfileremoval-unreferenceddays
	//
	UnreferencedDays *float64 `field:"optional" json:"unreferencedDays" yaml:"unreferencedDays"`
}

