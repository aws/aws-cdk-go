package awsworkspacesweb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkSettings.
// Experimental.
type INetworkSettingsRef interface {
	constructs.IConstruct
}

// The jsii proxy for INetworkSettingsRef
type jsiiProxy_INetworkSettingsRef struct {
	internal.Type__constructsIConstruct
}

