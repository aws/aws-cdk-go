package previewawsnetworkmanagerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.networkmanager@NetworkManagerSegmentUpdate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkManagerSegmentUpdate := awscdkmixinspreview.Events.NewNetworkManagerSegmentUpdate()
//
// Experimental.
type NetworkManagerSegmentUpdate interface {
}

// The jsii proxy struct for NetworkManagerSegmentUpdate
type jsiiProxy_NetworkManagerSegmentUpdate struct {
	_ byte // padding
}

// Experimental.
func NewNetworkManagerSegmentUpdate() NetworkManagerSegmentUpdate {
	_init_.Initialize()

	j := jsiiProxy_NetworkManagerSegmentUpdate{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerSegmentUpdate",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkManagerSegmentUpdate_Override(n NetworkManagerSegmentUpdate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerSegmentUpdate",
		nil, // no parameters
		n,
	)
}

// EventBridge event pattern for Network Manager Segment Update.
// Experimental.
func NetworkManagerSegmentUpdate_EventPattern(options *NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateNetworkManagerSegmentUpdate_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.NetworkManagerSegmentUpdate",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

