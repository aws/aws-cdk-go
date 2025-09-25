package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalyMonitor.
// Experimental.
type IAnomalyMonitorRef interface {
	constructs.IConstruct
	// A reference to a AnomalyMonitor resource.
	// Experimental.
	AnomalyMonitorRef() *AnomalyMonitorReference
}

// The jsii proxy for IAnomalyMonitorRef
type jsiiProxy_IAnomalyMonitorRef struct {
	internal.Type__constructsIConstruct
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

