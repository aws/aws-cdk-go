package interfacesawsresourceexplorer2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DefaultViewAssociation.
// Experimental.
type IDefaultViewAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DefaultViewAssociation resource.
	// Experimental.
	DefaultViewAssociationRef() *DefaultViewAssociationReference
}

// The jsii proxy for IDefaultViewAssociationRef
type jsiiProxy_IDefaultViewAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDefaultViewAssociationRef) DefaultViewAssociationRef() *DefaultViewAssociationReference {
	var returns *DefaultViewAssociationReference
	_jsii_.Get(
		j,
		"defaultViewAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDefaultViewAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDefaultViewAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

