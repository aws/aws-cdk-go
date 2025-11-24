package mixinsawscodecommit


// Information about code to be committed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeProperty := &CodeProperty{
//   	BranchName: jsii.String("branchName"),
//   	S3: &S3Property{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		ObjectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html
//
type CfnRepositoryPropsMixin_CodeProperty struct {
	// Optional.
	//
	// Specifies a branch name to be used as the default branch when importing code into a repository on initial creation. If this property is not set, the name *main* will be used for the default branch for the repository. Changes to this property are ignored after initial resource creation. We recommend using this parameter to set the name to *main* to align with the default behavior of CodeCommit unless another name is needed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-branchname
	//
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	//
	// Changes to this property are ignored after initial resource creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

