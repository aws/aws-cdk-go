package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslogs/internal"
)

// Experimental.
type ILogStream interface {
	awscdk.IResource
	// The name of this log stream.
	// Experimental.
	LogStreamName() *string
}

// The jsii proxy for ILogStream
type jsiiProxy_ILogStream struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ILogStream) LogStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logStreamName",
		&returns,
	)
	return returns
}

