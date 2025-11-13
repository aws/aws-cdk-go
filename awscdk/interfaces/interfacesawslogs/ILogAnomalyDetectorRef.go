package interfacesawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogAnomalyDetector.
// Experimental.
type ILogAnomalyDetectorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LogAnomalyDetector resource.
	// Experimental.
	LogAnomalyDetectorRef() *LogAnomalyDetectorReference
}

// The jsii proxy for ILogAnomalyDetectorRef
type jsiiProxy_ILogAnomalyDetectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ILogAnomalyDetectorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogAnomalyDetectorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

