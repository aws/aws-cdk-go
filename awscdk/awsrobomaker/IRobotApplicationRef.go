package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplication.
// Experimental.
type IRobotApplicationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RobotApplication resource.
	// Experimental.
	RobotApplicationRef() *RobotApplicationReference
}

// The jsii proxy for IRobotApplicationRef
type jsiiProxy_IRobotApplicationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRobotApplicationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

