package interfacesawssecurityagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TargetDomain.
// Experimental.
type ITargetDomainRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TargetDomain resource.
	// Experimental.
	TargetDomainRef() *TargetDomainReference
}

// The jsii proxy for ITargetDomainRef
type jsiiProxy_ITargetDomainRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITargetDomainRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITargetDomainRef) TargetDomainRef() *TargetDomainReference {
	var returns *TargetDomainReference
	_jsii_.Get(
		j,
		"targetDomainRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetDomainRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITargetDomainRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

