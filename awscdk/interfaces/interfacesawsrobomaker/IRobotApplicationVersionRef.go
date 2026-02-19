package interfacesawsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplicationVersion.
// Experimental.
type IRobotApplicationVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RobotApplicationVersion resource.
	// Experimental.
	RobotApplicationVersionRef() *RobotApplicationVersionReference
}

// The jsii proxy for IRobotApplicationVersionRef
type jsiiProxy_IRobotApplicationVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRobotApplicationVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRobotApplicationVersionRef) RobotApplicationVersionRef() *RobotApplicationVersionReference {
	var returns *RobotApplicationVersionReference
	_jsii_.Get(
		j,
		"robotApplicationVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRobotApplicationVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRobotApplicationVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

