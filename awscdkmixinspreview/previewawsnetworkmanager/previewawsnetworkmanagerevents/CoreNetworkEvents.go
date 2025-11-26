package previewawsnetworkmanagerevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager"
)

// EventBridge event patterns for CoreNetwork.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var coreNetworkRef ICoreNetworkRef
//
//   coreNetworkEvents := awscdkmixinspreview.Events.CoreNetworkEvents_FromCoreNetwork(coreNetworkRef)
//
// Experimental.
type CoreNetworkEvents interface {
	// EventBridge event pattern for CoreNetwork Network Manager Policy Update.
	// Experimental.
	NetworkManagerPolicyUpdatePattern(options *CoreNetworkEvents_NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps) *awsevents.EventPattern
	// EventBridge event pattern for CoreNetwork Network Manager Segment Update.
	// Experimental.
	NetworkManagerSegmentUpdatePattern(options *CoreNetworkEvents_NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps) *awsevents.EventPattern
}

// The jsii proxy struct for CoreNetworkEvents
type jsiiProxy_CoreNetworkEvents struct {
	_ byte // padding
}

// Create CoreNetworkEvents from a CoreNetwork reference.
// Experimental.
func CoreNetworkEvents_FromCoreNetwork(coreNetworkRef interfacesawsnetworkmanager.ICoreNetworkRef) CoreNetworkEvents {
	_init_.Initialize()

	if err := validateCoreNetworkEvents_FromCoreNetworkParameters(coreNetworkRef); err != nil {
		panic(err)
	}
	var returns CoreNetworkEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkmanager.events.CoreNetworkEvents",
		"fromCoreNetwork",
		[]interface{}{coreNetworkRef},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CoreNetworkEvents) NetworkManagerPolicyUpdatePattern(options *CoreNetworkEvents_NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps) *awsevents.EventPattern {
	if err := c.validateNetworkManagerPolicyUpdatePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"networkManagerPolicyUpdatePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CoreNetworkEvents) NetworkManagerSegmentUpdatePattern(options *CoreNetworkEvents_NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps) *awsevents.EventPattern {
	if err := c.validateNetworkManagerSegmentUpdatePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"networkManagerSegmentUpdatePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

