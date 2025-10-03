package awscodestar

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestar/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GitHubRepository.
// Experimental.
type IGitHubRepositoryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGitHubRepositoryRef
type jsiiProxy_IGitHubRepositoryRef struct {
	internal.Type__constructsIConstruct
}

