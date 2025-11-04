package awscodestar

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestar/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GitHubRepository.
// Experimental.
type IGitHubRepositoryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GitHubRepository resource.
	// Experimental.
	GitHubRepositoryRef() *GitHubRepositoryReference
}

// The jsii proxy for IGitHubRepositoryRef
type jsiiProxy_IGitHubRepositoryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IGitHubRepositoryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGitHubRepositoryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

