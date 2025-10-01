package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogStream.
// Experimental.
type ILogStreamRef interface {
	constructs.IConstruct
	// A reference to a LogStream resource.
	// Experimental.
	LogStreamRef() *LogStreamReference
}

// The jsii proxy for ILogStreamRef
type jsiiProxy_ILogStreamRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILogStreamRef) LogStreamRef() *LogStreamReference {
	var returns *LogStreamReference
	_jsii_.Get(
		j,
		"logStreamRef",
		&returns,
	)
	return returns
}

