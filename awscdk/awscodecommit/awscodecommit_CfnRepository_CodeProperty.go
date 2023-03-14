package awscodecommit


// Information about code to be committed.
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
//
//   	// the properties below are optional
//   	branchName: jsii.String("branchName"),
//   }
//
type CfnRepository_CodeProperty struct {
	// Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository.
	//
	// Changes to this property are ignored after initial resource creation.
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
	// Optional.
	//
	// Specifies a branch name to be used as the default branch when importing code into a repository on initial creation. If this property is not set, the name *main* will be used for the default branch for the repository. Changes to this property are ignored after initial resource creation. We recommend using this parameter to set the name to *main* to align with the default behavior of CodeCommit unless another name is needed.
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
}

