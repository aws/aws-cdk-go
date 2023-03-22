package awssagemaker


// Contains details regarding the file source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSourceProperty := &fileSourceProperty{
//   	s3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	contentDigest: jsii.String("contentDigest"),
//   	contentType: jsii.String("contentType"),
//   }
//
type CfnModelPackage_FileSourceProperty struct {
	// The Amazon S3 URI for the file source.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The digest of the file source.
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
	// The type of content stored in the file source.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

