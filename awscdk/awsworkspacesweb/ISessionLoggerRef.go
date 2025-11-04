package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SessionLogger.
// Experimental.
type ISessionLoggerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SessionLogger resource.
	// Experimental.
	SessionLoggerRef() *SessionLoggerReference
}

// The jsii proxy for ISessionLoggerRef
type jsiiProxy_ISessionLoggerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISessionLoggerRef) SessionLoggerRef() *SessionLoggerReference {
	var returns *SessionLoggerReference
	_jsii_.Get(
		j,
		"sessionLoggerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionLoggerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionLoggerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

