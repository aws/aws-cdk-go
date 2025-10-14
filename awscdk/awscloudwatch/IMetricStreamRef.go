package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MetricStream.
// Experimental.
type IMetricStreamRef interface {
	constructs.IConstruct
	// A reference to a MetricStream resource.
	// Experimental.
	MetricStreamRef() *MetricStreamReference
}

// The jsii proxy for IMetricStreamRef
type jsiiProxy_IMetricStreamRef struct {
	internal.Type__constructsIConstruct
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

