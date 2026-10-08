package interfacesawsram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrincipalAssociation.
// Experimental.
type IPrincipalAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PrincipalAssociation resource.
	// Experimental.
	PrincipalAssociationRef() *PrincipalAssociationReference
}

// The jsii proxy for IPrincipalAssociationRef
type jsiiProxy_IPrincipalAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPrincipalAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPrincipalAssociationRef) PrincipalAssociationRef() *PrincipalAssociationReference {
	var returns *PrincipalAssociationReference
	_jsii_.Get(
		j,
		"principalAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrincipalAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrincipalAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

