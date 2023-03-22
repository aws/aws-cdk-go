package awscodestar


// The `S3` property type specifies information about the Amazon S3 bucket that contains the code to be committed to the new repository.
//
// `S3` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3Property := &s3Property{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnGitHubRepository_S3Property struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 object key or file name for the ZIP file.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

