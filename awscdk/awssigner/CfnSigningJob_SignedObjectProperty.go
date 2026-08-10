package awssigner


// The S3 location of the signed code image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signedObjectProperty := &SignedObjectProperty{
//   	S3: &S3SignedObjectProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-signedobject.html
//
type CfnSigningJob_SignedObjectProperty struct {
	// The Amazon S3 bucket name and key where Signer saved the signed code image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-signer-signingjob-signedobject.html#cfn-signer-signingjob-signedobject-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

