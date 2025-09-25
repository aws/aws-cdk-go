package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Logging.
// Experimental.
type ILoggingRef interface {
	constructs.IConstruct
	// A reference to a Logging resource.
	// Experimental.
	LoggingRef() *LoggingReference
}

// The jsii proxy for ILoggingRef
type jsiiProxy_ILoggingRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILoggingRef) LoggingRef() *LoggingReference {
	var returns *LoggingReference
	_jsii_.Get(
		j,
		"loggingRef",
		&returns,
	)
	return returns
}

