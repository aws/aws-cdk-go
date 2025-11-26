package previewawsnetworkmanagerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.networkmanager@NetworkManagerPolicyUpdate event types for CoreNetwork.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkManagerPolicyUpdate := #error#.NewNetworkManagerPolicyUpdate()
//
// Experimental.
type CoreNetworkEvents_NetworkManagerPolicyUpdate interface {
}

// The jsii proxy struct for CoreNetworkEvents_NetworkManagerPolicyUpdate
type jsiiProxy_CoreNetworkEvents_NetworkManagerPolicyUpdate struct {
	_ byte // padding
}

// Experimental.
func NewCoreNetworkEvents_NetworkManagerPolicyUpdate() CoreNetworkEvents_NetworkManagerPolicyUpdate {
	_init_.Initialize()

	j := jsiiProxy_CoreNetworkEvents_NetworkManagerPolicyUpdate{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerPolicyUpdate",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCoreNetworkEvents_NetworkManagerPolicyUpdate_Override(c CoreNetworkEvents_NetworkManagerPolicyUpdate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerPolicyUpdate",
		nil, // no parameters
		c,
	)
}

