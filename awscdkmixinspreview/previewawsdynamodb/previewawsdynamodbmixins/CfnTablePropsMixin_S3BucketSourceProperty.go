package previewawsdynamodbmixins


// The S3 bucket that is being imported from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3BucketSourceProperty := &S3BucketSourceProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3BucketOwner: jsii.String("s3BucketOwner"),
//   	S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-s3bucketsource.html
//
type CfnTablePropsMixin_S3BucketSourceProperty struct {
	// The S3 bucket that is being imported from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-s3bucketsource.html#cfn-dynamodb-table-s3bucketsource-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The account number of the S3 bucket that is being imported from.
	//
	// If the bucket is owned by the requester this is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-s3bucketsource.html#cfn-dynamodb-table-s3bucketsource-s3bucketowner
	//
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// The key prefix shared by all S3 Objects that are being imported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-s3bucketsource.html#cfn-dynamodb-table-s3bucketsource-s3keyprefix
	//
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

