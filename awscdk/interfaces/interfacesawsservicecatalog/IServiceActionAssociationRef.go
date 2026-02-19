package interfacesawsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceActionAssociation.
// Experimental.
type IServiceActionAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServiceActionAssociation resource.
	// Experimental.
	ServiceActionAssociationRef() *ServiceActionAssociationReference
}

// The jsii proxy for IServiceActionAssociationRef
type jsiiProxy_IServiceActionAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServiceActionAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IServiceActionAssociationRef) ServiceActionAssociationRef() *ServiceActionAssociationReference {
	var returns *ServiceActionAssociationReference
	_jsii_.Get(
		j,
		"serviceActionAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceActionAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceActionAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

