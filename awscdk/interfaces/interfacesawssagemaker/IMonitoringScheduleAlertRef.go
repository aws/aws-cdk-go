package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MonitoringScheduleAlert.
// Experimental.
type IMonitoringScheduleAlertRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MonitoringScheduleAlert resource.
	// Experimental.
	MonitoringScheduleAlertRef() *MonitoringScheduleAlertReference
}

// The jsii proxy for IMonitoringScheduleAlertRef
type jsiiProxy_IMonitoringScheduleAlertRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMonitoringScheduleAlertRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IMonitoringScheduleAlertRef) MonitoringScheduleAlertRef() *MonitoringScheduleAlertReference {
	var returns *MonitoringScheduleAlertReference
	_jsii_.Get(
		j,
		"monitoringScheduleAlertRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitoringScheduleAlertRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitoringScheduleAlertRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

