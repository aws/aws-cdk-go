package awsinternetmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinternetmonitor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Monitor.
// Experimental.
type IMonitorRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Monitor resource.
	// Experimental.
	MonitorRef() *MonitorReference
}

// The jsii proxy for IMonitorRef
type jsiiProxy_IMonitorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IMonitorRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

