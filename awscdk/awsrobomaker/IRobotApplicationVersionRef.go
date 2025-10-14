package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplicationVersion.
// Experimental.
type IRobotApplicationVersionRef interface {
	constructs.IConstruct
	// A reference to a RobotApplicationVersion resource.
	// Experimental.
	RobotApplicationVersionRef() *RobotApplicationVersionReference
}

// The jsii proxy for IRobotApplicationVersionRef
type jsiiProxy_IRobotApplicationVersionRef struct {
	internal.Type__constructsIConstruct
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

