package awscodestarconnections

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryLink.
// Experimental.
type IRepositoryLinkRef interface {
	constructs.IConstruct
	// A reference to a RepositoryLink resource.
	// Experimental.
	RepositoryLinkRef() *RepositoryLinkReference
}

// The jsii proxy for IRepositoryLinkRef
type jsiiProxy_IRepositoryLinkRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRepositoryLinkRef) RepositoryLinkRef() *RepositoryLinkReference {
	var returns *RepositoryLinkReference
	_jsii_.Get(
		j,
		"repositoryLinkRef",
		&returns,
	)
	return returns
}

