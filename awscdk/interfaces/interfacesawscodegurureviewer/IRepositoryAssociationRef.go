package interfacesawscodegurureviewer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodegurureviewer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryAssociation.
// Experimental.
type IRepositoryAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RepositoryAssociation resource.
	// Experimental.
	RepositoryAssociationRef() *RepositoryAssociationReference
}

// The jsii proxy for IRepositoryAssociationRef
type jsiiProxy_IRepositoryAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRepositoryAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IRepositoryAssociationRef) RepositoryAssociationRef() *RepositoryAssociationReference {
	var returns *RepositoryAssociationReference
	_jsii_.Get(
		j,
		"repositoryAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepositoryAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepositoryAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

