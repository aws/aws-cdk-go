package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContextStoreAssociation.
// Experimental.
type IContextStoreAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContextStoreAssociation resource.
	// Experimental.
	ContextStoreAssociationRef() *ContextStoreAssociationReference
}

// The jsii proxy for IContextStoreAssociationRef
type jsiiProxy_IContextStoreAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IContextStoreAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IContextStoreAssociationRef) ContextStoreAssociationRef() *ContextStoreAssociationReference {
	var returns *ContextStoreAssociationReference
	_jsii_.Get(
		j,
		"contextStoreAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextStoreAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextStoreAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

