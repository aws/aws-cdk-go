package awsssmguiconnect


// The S3 bucket where RDP connection recordings are stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3BucketProperty := &S3BucketProperty{
//   	BucketName: jsii.String("bucketName"),
//   	BucketOwner: jsii.String("bucketOwner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-s3bucket.html
//
type CfnPreferences_S3BucketProperty struct {
	// The name of the S3 bucket where RDP connection recordings are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-s3bucket.html#cfn-ssmguiconnect-preferences-s3bucket-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The AWS account number that owns the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmguiconnect-preferences-s3bucket.html#cfn-ssmguiconnect-preferences-s3bucket-bucketowner
	//
	BucketOwner *string `field:"required" json:"bucketOwner" yaml:"bucketOwner"`
}

