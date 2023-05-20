package awscdkcodestaralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcodestaralpha/v2/internal"
)

// GitHubRepository resource interface.
// Experimental.
type IGitHubRepository interface {
	awscdk.IResource
	// the repository owner.
	// Experimental.
	Owner() *string
	// the repository name.
	// Experimental.
	Repo() *string
}

// The jsii proxy for IGitHubRepository
type jsiiProxy_IGitHubRepository struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IGitHubRepository) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGitHubRepository) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

