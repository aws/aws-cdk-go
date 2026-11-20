package awskinesisanalyticsv2


// An Amazon S3 prefix location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3PrefixLocationProperty := &S3PrefixLocationProperty{
//   	BucketArn: jsii.String("bucketArn"),
//   	KeyPrefix: jsii.String("keyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-s3prefixlocation.html
//
type CfnSnapshotPropsMixin_S3PrefixLocationProperty struct {
	// The ARN of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-s3prefixlocation.html#cfn-kinesisanalyticsv2-snapshot-s3prefixlocation-bucketarn
	//
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The S3 key prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-snapshot-s3prefixlocation.html#cfn-kinesisanalyticsv2-snapshot-s3prefixlocation-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

