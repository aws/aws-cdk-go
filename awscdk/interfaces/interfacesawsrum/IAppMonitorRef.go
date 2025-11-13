package interfacesawsrum

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrum/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppMonitor.
// Experimental.
type IAppMonitorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AppMonitor resource.
	// Experimental.
	AppMonitorRef() *AppMonitorReference
}

// The jsii proxy for IAppMonitorRef
type jsiiProxy_IAppMonitorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAppMonitorRef) AppMonitorRef() *AppMonitorReference {
	var returns *AppMonitorReference
	_jsii_.Get(
		j,
		"appMonitorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppMonitorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppMonitorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

