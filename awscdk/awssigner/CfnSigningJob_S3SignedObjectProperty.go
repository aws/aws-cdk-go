package awssigner


// The Amazon S3 bucket name and key where Signer saved the signed code image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SignedObjectProperty := &S3SignedObjectProperty{
//   	BucketName: jsii.String("bucketName"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3signedobject.html
//
type CfnSigningJob_S3SignedObjectProperty struct {
	// Name of the S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3signedobject.html#cfn-signer-signingjob-s3signedobject-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Key name that uniquely identifies a signed code image in the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-s3signedobject.html#cfn-signer-signingjob-s3signedobject-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}

