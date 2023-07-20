package awssagemaker


// Contains details regarding the file source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSourceProperty := &FileSourceProperty{
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	ContentDigest: jsii.String("contentDigest"),
//   	ContentType: jsii.String("contentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-filesource.html
//
type CfnModelPackage_FileSourceProperty struct {
	// The Amazon S3 URI for the file source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-filesource.html#cfn-sagemaker-modelpackage-filesource-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The digest of the file source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-filesource.html#cfn-sagemaker-modelpackage-filesource-contentdigest
	//
	ContentDigest *string `field:"optional" json:"contentDigest" yaml:"contentDigest"`
	// The type of content stored in the file source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-filesource.html#cfn-sagemaker-modelpackage-filesource-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

