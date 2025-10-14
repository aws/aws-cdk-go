package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Robot.
// Experimental.
type IRobotRef interface {
	constructs.IConstruct
	// A reference to a Robot resource.
	// Experimental.
	RobotRef() *RobotReference
}

// The jsii proxy for IRobotRef
type jsiiProxy_IRobotRef struct {
	internal.Type__constructsIConstruct
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

