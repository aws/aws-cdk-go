package interfacesawscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdNamespaceAssociation.
// Experimental.
type IIdNamespaceAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdNamespaceAssociation resource.
	// Experimental.
	IdNamespaceAssociationRef() *IdNamespaceAssociationReference
}

// The jsii proxy for IIdNamespaceAssociationRef
type jsiiProxy_IIdNamespaceAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIdNamespaceAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIdNamespaceAssociationRef) IdNamespaceAssociationRef() *IdNamespaceAssociationReference {
	var returns *IdNamespaceAssociationReference
	_jsii_.Get(
		j,
		"idNamespaceAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdNamespaceAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdNamespaceAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

