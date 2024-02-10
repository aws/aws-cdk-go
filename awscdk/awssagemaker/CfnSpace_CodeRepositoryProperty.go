package awssagemaker


// A Git repository that SageMaker automatically displays to users for cloning in the JupyterServer application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeRepositoryProperty := &CodeRepositoryProperty{
//   	RepositoryUrl: jsii.String("repositoryUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-coderepository.html
//
type CfnSpace_CodeRepositoryProperty struct {
	// The URL of the Git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-coderepository.html#cfn-sagemaker-space-coderepository-repositoryurl
	//
	RepositoryUrl *string `field:"required" json:"repositoryUrl" yaml:"repositoryUrl"`
}

