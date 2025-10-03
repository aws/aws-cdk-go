package awssupportapp

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssupportapp/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SlackWorkspaceConfiguration.
// Experimental.
type ISlackWorkspaceConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISlackWorkspaceConfigurationRef
type jsiiProxy_ISlackWorkspaceConfigurationRef struct {
	internal.Type__constructsIConstruct
}

