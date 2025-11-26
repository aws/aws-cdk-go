package previewawscodestarmixins


// The `S3` property type specifies information about the Amazon S3 bucket that contains the code to be committed to the new repository.
//
// `S3` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3Property := &S3Property{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   	ObjectVersion: jsii.String("objectVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestar-githubrepository-s3.html
//
type CfnGitHubRepositoryPropsMixin_S3Property struct {
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestar-githubrepository-s3.html#cfn-codestar-githubrepository-s3-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 object key or file name for the ZIP file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestar-githubrepository-s3.html#cfn-codestar-githubrepository-s3-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestar-githubrepository-s3.html#cfn-codestar-githubrepository-s3-objectversion
	//
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

