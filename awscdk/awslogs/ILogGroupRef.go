package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogGroup.
// Experimental.
type ILogGroupRef interface {
	constructs.IConstruct
	// A reference to a LogGroup resource.
	// Experimental.
	LogGroupRef() *LogGroupReference
}

// The jsii proxy for ILogGroupRef
type jsiiProxy_ILogGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILogGroupRef) LogGroupRef() *LogGroupReference {
	var returns *LogGroupReference
	_jsii_.Get(
		j,
		"logGroupRef",
		&returns,
	)
	return returns
}

