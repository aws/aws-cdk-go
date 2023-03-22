package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCodeRepository`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeRepositoryProps := &cfnCodeRepositoryProps{
//   	gitConfig: &gitConfigProperty{
//   		repositoryUrl: jsii.String("repositoryUrl"),
//
//   		// the properties below are optional
//   		branch: jsii.String("branch"),
//   		secretArn: jsii.String("secretArn"),
//   	},
//
//   	// the properties below are optional
//   	codeRepositoryName: jsii.String("codeRepositoryName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCodeRepositoryProps struct {
	// Configuration details for the Git repository, including the URL where it is located and the ARN of the AWS Secrets Manager secret that contains the credentials used to access the repository.
	GitConfig interface{} `field:"required" json:"gitConfig" yaml:"gitConfig"`
	// The name of the Git repository.
	CodeRepositoryName *string `field:"optional" json:"codeRepositoryName" yaml:"codeRepositoryName"`
	// List of tags for Code Repository.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

