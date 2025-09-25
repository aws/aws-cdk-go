package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogAnomalyDetector.
// Experimental.
type ILogAnomalyDetectorRef interface {
	constructs.IConstruct
	// A reference to a LogAnomalyDetector resource.
	// Experimental.
	LogAnomalyDetectorRef() *LogAnomalyDetectorReference
}

// The jsii proxy for ILogAnomalyDetectorRef
type jsiiProxy_ILogAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILogAnomalyDetectorRef) LogAnomalyDetectorRef() *LogAnomalyDetectorReference {
	var returns *LogAnomalyDetectorReference
	_jsii_.Get(
		j,
		"logAnomalyDetectorRef",
		&returns,
	)
	return returns
}

