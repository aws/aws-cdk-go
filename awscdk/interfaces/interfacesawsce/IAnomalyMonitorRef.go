package interfacesawsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyMonitor.
// Experimental.
type IAnomalyMonitorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AnomalyMonitor resource.
	// Experimental.
	AnomalyMonitorRef() *AnomalyMonitorReference
}

// The jsii proxy for IAnomalyMonitorRef
type jsiiProxy_IAnomalyMonitorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAnomalyMonitorRef) AnomalyMonitorRef() *AnomalyMonitorReference {
	var returns *AnomalyMonitorReference
	_jsii_.Get(
		j,
		"anomalyMonitorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyMonitorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyMonitorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

