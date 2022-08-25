// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
)

// Properties for a CodeCommit source code provider.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repository := codecommit.NewRepository(this, jsii.String("Repo"), &repositoryProps{
//   	repositoryName: jsii.String("my-repo"),
//   })
//
//   amplifyApp := amplify.NewApp(this, jsii.String("App"), &appProps{
//   	sourceCodeProvider: amplify.NewCodeCommitSourceCodeProvider(&codeCommitSourceCodeProviderProps{
//   		repository: repository,
//   	}),
//   })
//
// Experimental.
type CodeCommitSourceCodeProviderProps struct {
	// The CodeCommit repository.
	// Experimental.
	Repository awscodecommit.IRepository `field:"required" json:"repository" yaml:"repository"`
}

