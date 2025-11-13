package interfacesawsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Robot.
// Experimental.
type IRobotRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Robot resource.
	// Experimental.
	RobotRef() *RobotReference
}

// The jsii proxy for IRobotRef
type jsiiProxy_IRobotRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRobotRef) RobotRef() *RobotReference {
	var returns *RobotReference
	_jsii_.Get(
		j,
		"robotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRobotRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRobotRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

