package previewawsnetworkmanagerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.networkmanager@NetworkManagerSegmentUpdate event types for CoreNetwork.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkManagerSegmentUpdate := #error#.NewNetworkManagerSegmentUpdate()
//
// Experimental.
type CoreNetworkEvents_NetworkManagerSegmentUpdate interface {
}

// The jsii proxy struct for CoreNetworkEvents_NetworkManagerSegmentUpdate
type jsiiProxy_CoreNetworkEvents_NetworkManagerSegmentUpdate struct {
	_ byte // padding
}

// Experimental.
func NewCoreNetworkEvents_NetworkManagerSegmentUpdate() CoreNetworkEvents_NetworkManagerSegmentUpdate {
	_init_.Initialize()

	j := jsiiProxy_CoreNetworkEvents_NetworkManagerSegmentUpdate{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerSegmentUpdate",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCoreNetworkEvents_NetworkManagerSegmentUpdate_Override(c CoreNetworkEvents_NetworkManagerSegmentUpdate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents.NetworkManagerSegmentUpdate",
		nil, // no parameters
		c,
	)
}

