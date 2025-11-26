package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCodeRepositoryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeRepositoryMixinProps := &CfnCodeRepositoryMixinProps{
//   	CodeRepositoryName: jsii.String("codeRepositoryName"),
//   	GitConfig: &GitConfigProperty{
//   		Branch: jsii.String("branch"),
//   		RepositoryUrl: jsii.String("repositoryUrl"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-coderepository.html
//
type CfnCodeRepositoryMixinProps struct {
	// The name of the Git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-coderepository.html#cfn-sagemaker-coderepository-coderepositoryname
	//
	CodeRepositoryName *string `field:"optional" json:"codeRepositoryName" yaml:"codeRepositoryName"`
	// Configuration details for the Git repository, including the URL where it is located and the ARN of the AWS Secrets Manager secret that contains the credentials used to access the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-coderepository.html#cfn-sagemaker-coderepository-gitconfig
	//
	GitConfig interface{} `field:"optional" json:"gitConfig" yaml:"gitConfig"`
	// List of tags for Code Repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-coderepository.html#cfn-sagemaker-coderepository-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

