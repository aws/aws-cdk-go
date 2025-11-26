package previewawscodestarmixins


// The `Code` property type specifies information about code to be committed.
//
// `Code` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeProperty := &CodeProperty{
//   	S3: &S3Property{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestar-githubrepository-code.html
//
type CfnGitHubRepositoryPropsMixin_CodeProperty struct {
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codestar-githubrepository-code.html#cfn-codestar-githubrepository-code-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

