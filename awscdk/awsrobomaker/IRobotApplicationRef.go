package awsrobomaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrobomaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RobotApplication.
// Experimental.
type IRobotApplicationRef interface {
	constructs.IConstruct
	// A reference to a RobotApplication resource.
	// Experimental.
	RobotApplicationRef() *RobotApplicationReference
}

// The jsii proxy for IRobotApplicationRef
type jsiiProxy_IRobotApplicationRef struct {
	internal.Type__constructsIConstruct
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

