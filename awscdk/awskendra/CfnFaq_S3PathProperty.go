package awskendra


// Information required to find a specific file in an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3PathProperty := &S3PathProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-faq-s3path.html
//
type CfnFaq_S3PathProperty struct {
	// The name of the S3 bucket that contains the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-faq-s3path.html#cfn-kendra-faq-s3path-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-faq-s3path.html#cfn-kendra-faq-s3path-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
}

