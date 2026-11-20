package interfacesawsnetworkflowmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkflowmonitor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PerformanceMonitor.
// Experimental.
type IPerformanceMonitorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PerformanceMonitor resource.
	// Experimental.
	PerformanceMonitorRef() *PerformanceMonitorReference
}

// The jsii proxy for IPerformanceMonitorRef
type jsiiProxy_IPerformanceMonitorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPerformanceMonitorRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPerformanceMonitorRef) PerformanceMonitorRef() *PerformanceMonitorReference {
	var returns *PerformanceMonitorReference
	_jsii_.Get(
		j,
		"performanceMonitorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformanceMonitorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPerformanceMonitorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

