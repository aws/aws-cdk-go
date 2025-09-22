package awssupportapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SlackWorkspaceConfiguration.
// Experimental.
type ISlackWorkspaceConfigurationRef interface {
	constructs.IConstruct
	// A reference to a SlackWorkspaceConfiguration resource.
	// Experimental.
	SlackWorkspaceConfigurationRef() *SlackWorkspaceConfigurationReference
}

// The jsii proxy for ISlackWorkspaceConfigurationRef
type jsiiProxy_ISlackWorkspaceConfigurationRef struct {
	internal.Type__constructsIConstruct
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

