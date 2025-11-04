package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MonitoringSchedule.
// Experimental.
type IMonitoringScheduleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MonitoringSchedule resource.
	// Experimental.
	MonitoringScheduleRef() *MonitoringScheduleReference
}

// The jsii proxy for IMonitoringScheduleRef
type jsiiProxy_IMonitoringScheduleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMonitoringScheduleRef) MonitoringScheduleRef() *MonitoringScheduleReference {
	var returns *MonitoringScheduleReference
	_jsii_.Get(
		j,
		"monitoringScheduleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitoringScheduleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitoringScheduleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

