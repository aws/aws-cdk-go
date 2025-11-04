package awsdatabrew

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Schedule.
// Experimental.
type IScheduleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Schedule resource.
	// Experimental.
	ScheduleRef() *ScheduleReference
}

// The jsii proxy for IScheduleRef
type jsiiProxy_IScheduleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IScheduleRef) ScheduleRef() *ScheduleReference {
	var returns *ScheduleReference
	_jsii_.Get(
		j,
		"scheduleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScheduleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

