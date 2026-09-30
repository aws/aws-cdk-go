package interfacesawsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QEV2IdcApplication.
// Experimental.
type IQEV2IdcApplicationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QEV2IdcApplication resource.
	// Experimental.
	Qev2IdcApplicationRef() *QEV2IdcApplicationReference
}

// The jsii proxy for IQEV2IdcApplicationRef
type jsiiProxy_IQEV2IdcApplicationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IQEV2IdcApplicationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IQEV2IdcApplicationRef) Qev2IdcApplicationRef() *QEV2IdcApplicationReference {
	var returns *QEV2IdcApplicationReference
	_jsii_.Get(
		j,
		"qev2IdcApplicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQEV2IdcApplicationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQEV2IdcApplicationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

