package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplicationVersion.
// Experimental.
type IRobotApplicationVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RobotApplicationVersion resource.
	// Experimental.
	RobotApplicationVersionRef() *RobotApplicationVersionReference
}

// The jsii proxy for IRobotApplicationVersionRef
type jsiiProxy_IRobotApplicationVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRobotApplicationVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

