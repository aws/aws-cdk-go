package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentProfile.
// Experimental.
type IEnvironmentProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnvironmentProfile resource.
	// Experimental.
	EnvironmentProfileRef() *EnvironmentProfileReference
}

// The jsii proxy for IEnvironmentProfileRef
type jsiiProxy_IEnvironmentProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEnvironmentProfileRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEnvironmentProfileRef) EnvironmentProfileRef() *EnvironmentProfileReference {
	var returns *EnvironmentProfileReference
	_jsii_.Get(
		j,
		"environmentProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

