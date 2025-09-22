package awsworkspacesweb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesweb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkSettings.
// Experimental.
type INetworkSettingsRef interface {
	constructs.IConstruct
	// A reference to a NetworkSettings resource.
	// Experimental.
	NetworkSettingsRef() *NetworkSettingsReference
}

// The jsii proxy for INetworkSettingsRef
type jsiiProxy_INetworkSettingsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkSettingsRef) NetworkSettingsRef() *NetworkSettingsReference {
	var returns *NetworkSettingsReference
	_jsii_.Get(
		j,
		"networkSettingsRef",
		&returns,
	)
	return returns
}

