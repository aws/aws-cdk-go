package awscodestar


// The `Code` property type specifies information about code to be committed.
//
// `Code` is a property of the `AWS::CodeStar::GitHubRepository` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeProperty := &codeProperty{
//   	s3: &s3Property{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type CfnGitHubRepository_CodeProperty struct {
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

