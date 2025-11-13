package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workteam.
// Experimental.
type IWorkteamRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Workteam resource.
	// Experimental.
	WorkteamRef() *WorkteamReference
}

// The jsii proxy for IWorkteamRef
type jsiiProxy_IWorkteamRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IWorkteamRef) WorkteamRef() *WorkteamReference {
	var returns *WorkteamReference
	_jsii_.Get(
		j,
		"workteamRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkteamRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkteamRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

