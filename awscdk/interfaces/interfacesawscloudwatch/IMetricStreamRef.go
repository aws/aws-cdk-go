package interfacesawscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MetricStream.
// Experimental.
type IMetricStreamRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MetricStream resource.
	// Experimental.
	MetricStreamRef() *MetricStreamReference
}

// The jsii proxy for IMetricStreamRef
type jsiiProxy_IMetricStreamRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMetricStreamRef) MetricStreamRef() *MetricStreamReference {
	var returns *MetricStreamReference
	_jsii_.Get(
		j,
		"metricStreamRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMetricStreamRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMetricStreamRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

