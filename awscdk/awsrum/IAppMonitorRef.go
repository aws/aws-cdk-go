package awsrum

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrum/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppMonitor.
// Experimental.
type IAppMonitorRef interface {
	constructs.IConstruct
	// A reference to a AppMonitor resource.
	// Experimental.
	AppMonitorRef() *AppMonitorReference
}

// The jsii proxy for IAppMonitorRef
type jsiiProxy_IAppMonitorRef struct {
	internal.Type__constructsIConstruct
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

