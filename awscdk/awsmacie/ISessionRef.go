package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Session.
// Experimental.
type ISessionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Session resource.
	// Experimental.
	SessionRef() *SessionReference
}

// The jsii proxy for ISessionRef
type jsiiProxy_ISessionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISessionRef) SessionRef() *SessionReference {
	var returns *SessionReference
	_jsii_.Get(
		j,
		"sessionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

