package awscodestar

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestar/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GitHubRepository.
// Experimental.
type IGitHubRepositoryRef interface {
	constructs.IConstruct
	// A reference to a GitHubRepository resource.
	// Experimental.
	GitHubRepositoryRef() *GitHubRepositoryReference
}

// The jsii proxy for IGitHubRepositoryRef
type jsiiProxy_IGitHubRepositoryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGitHubRepositoryRef) GitHubRepositoryRef() *GitHubRepositoryReference {
	var returns *GitHubRepositoryReference
	_jsii_.Get(
		j,
		"gitHubRepositoryRef",
		&returns,
	)
	return returns
}

