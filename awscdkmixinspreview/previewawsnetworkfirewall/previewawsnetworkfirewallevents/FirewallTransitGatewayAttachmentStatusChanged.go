package previewawsnetworkfirewallevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.networkfirewall@FirewallTransitGatewayAttachmentStatusChanged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firewallTransitGatewayAttachmentStatusChanged := awscdkmixinspreview.Events.NewFirewallTransitGatewayAttachmentStatusChanged()
//
// Experimental.
type FirewallTransitGatewayAttachmentStatusChanged interface {
}

// The jsii proxy struct for FirewallTransitGatewayAttachmentStatusChanged
type jsiiProxy_FirewallTransitGatewayAttachmentStatusChanged struct {
	_ byte // padding
}

// Experimental.
func NewFirewallTransitGatewayAttachmentStatusChanged() FirewallTransitGatewayAttachmentStatusChanged {
	_init_.Initialize()

	j := jsiiProxy_FirewallTransitGatewayAttachmentStatusChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallTransitGatewayAttachmentStatusChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallTransitGatewayAttachmentStatusChanged_Override(f FirewallTransitGatewayAttachmentStatusChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallTransitGatewayAttachmentStatusChanged",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for Firewall Transit Gateway Attachment Status Changed.
// Experimental.
func FirewallTransitGatewayAttachmentStatusChanged_FirewallTransitGatewayAttachmentStatusChangedPattern(options *FirewallTransitGatewayAttachmentStatusChanged_FirewallTransitGatewayAttachmentStatusChangedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFirewallTransitGatewayAttachmentStatusChanged_FirewallTransitGatewayAttachmentStatusChangedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallTransitGatewayAttachmentStatusChanged",
		"firewallTransitGatewayAttachmentStatusChangedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

