package interfacesawsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainVerification.
// Experimental.
type IDomainVerificationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainVerification resource.
	// Experimental.
	DomainVerificationRef() *DomainVerificationReference
}

// The jsii proxy for IDomainVerificationRef
type jsiiProxy_IDomainVerificationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDomainVerificationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDomainVerificationRef) DomainVerificationRef() *DomainVerificationReference {
	var returns *DomainVerificationReference
	_jsii_.Get(
		j,
		"domainVerificationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainVerificationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainVerificationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

