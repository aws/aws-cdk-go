package interfacesawsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplication.
// Experimental.
type IRobotApplicationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RobotApplication resource.
	// Experimental.
	RobotApplicationRef() *RobotApplicationReference
}

// The jsii proxy for IRobotApplicationRef
type jsiiProxy_IRobotApplicationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRobotApplicationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRobotApplicationRef) RobotApplicationRef() *RobotApplicationReference {
	var returns *RobotApplicationReference
	_jsii_.Get(
		j,
		"robotApplicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRobotApplicationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRobotApplicationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

