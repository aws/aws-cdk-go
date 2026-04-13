package previewawsnetworkmanagerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.networkmanager@NetworkManagerPolicyUpdate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkManagerPolicyUpdate := awscdkmixinspreview.Events.NewNetworkManagerPolicyUpdate()
//
// Experimental.
type NetworkManagerPolicyUpdate interface {
}

// The jsii proxy struct for NetworkManagerPolicyUpdate
type jsiiProxy_NetworkManagerPolicyUpdate struct {
	_ byte // padding
}

// Experimental.
func NewNetworkManagerPolicyUpdate() NetworkManagerPolicyUpdate {
	_init_.Initialize()

	j := jsiiProxy_NetworkManagerPolicyUpdate{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerPolicyUpdate",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkManagerPolicyUpdate_Override(n NetworkManagerPolicyUpdate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerPolicyUpdate",
		nil, // no parameters
		n,
	)
}

// EventBridge event pattern for Network Manager Policy Update.
// Experimental.
func NetworkManagerPolicyUpdate_EventPattern(options *NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateNetworkManagerPolicyUpdate_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerPolicyUpdate",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

