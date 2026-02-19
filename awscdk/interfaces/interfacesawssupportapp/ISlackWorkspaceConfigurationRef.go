package interfacesawssupportapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SlackWorkspaceConfiguration.
// Experimental.
type ISlackWorkspaceConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SlackWorkspaceConfiguration resource.
	// Experimental.
	SlackWorkspaceConfigurationRef() *SlackWorkspaceConfigurationReference
}

// The jsii proxy for ISlackWorkspaceConfigurationRef
type jsiiProxy_ISlackWorkspaceConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISlackWorkspaceConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISlackWorkspaceConfigurationRef) SlackWorkspaceConfigurationRef() *SlackWorkspaceConfigurationReference {
	var returns *SlackWorkspaceConfigurationReference
	_jsii_.Get(
		j,
		"slackWorkspaceConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackWorkspaceConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISlackWorkspaceConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

