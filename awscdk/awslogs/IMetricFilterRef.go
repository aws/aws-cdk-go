package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MetricFilter.
// Experimental.
type IMetricFilterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MetricFilter resource.
	// Experimental.
	MetricFilterRef() *MetricFilterReference
}

// The jsii proxy for IMetricFilterRef
type jsiiProxy_IMetricFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMetricFilterRef) MetricFilterRef() *MetricFilterReference {
	var returns *MetricFilterReference
	_jsii_.Get(
		j,
		"metricFilterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMetricFilterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMetricFilterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

