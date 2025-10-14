package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Monitor.
// Experimental.
type IMonitorRef interface {
	constructs.IConstruct
	// A reference to a Monitor resource.
	// Experimental.
	MonitorRef() *MonitorReference
}

// The jsii proxy for IMonitorRef
type jsiiProxy_IMonitorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMonitorRef) MonitorRef() *MonitorReference {
	var returns *MonitorReference
	_jsii_.Get(
		j,
		"monitorRef",
		&returns,
	)
	return returns
}

