package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogStream.
// Experimental.
type ILogStreamRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LogStream resource.
	// Experimental.
	LogStreamRef() *LogStreamReference
}

// The jsii proxy for ILogStreamRef
type jsiiProxy_ILogStreamRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ILogStreamRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogStreamRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

