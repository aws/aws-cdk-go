package interfacesawscodestarconnections

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryLink.
// Experimental.
type IRepositoryLinkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RepositoryLink resource.
	// Experimental.
	RepositoryLinkRef() *RepositoryLinkReference
}

// The jsii proxy for IRepositoryLinkRef
type jsiiProxy_IRepositoryLinkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IRepositoryLinkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepositoryLinkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

