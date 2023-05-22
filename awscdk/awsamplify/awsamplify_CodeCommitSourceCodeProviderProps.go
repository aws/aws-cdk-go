package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
)

// Properties for a CodeCommit source code provider.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repository := codecommit.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	RepositoryName: jsii.String("my-repo"),
//   })
//
//   amplifyApp := amplify.NewApp(this, jsii.String("App"), &AppProps{
//   	SourceCodeProvider: amplify.NewCodeCommitSourceCodeProvider(&CodeCommitSourceCodeProviderProps{
//   		Repository: *Repository,
//   	}),
//   })
//
// Experimental.
type CodeCommitSourceCodeProviderProps struct {
	// The CodeCommit repository.
	// Experimental.
	Repository awscodecommit.IRepository `field:"required" json:"repository" yaml:"repository"`
}

