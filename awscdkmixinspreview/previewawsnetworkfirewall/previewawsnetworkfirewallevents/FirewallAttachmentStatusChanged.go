package previewawsnetworkfirewallevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.networkfirewall@FirewallAttachmentStatusChanged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firewallAttachmentStatusChanged := awscdkmixinspreview.Events.NewFirewallAttachmentStatusChanged()
//
// Experimental.
type FirewallAttachmentStatusChanged interface {
}

// The jsii proxy struct for FirewallAttachmentStatusChanged
type jsiiProxy_FirewallAttachmentStatusChanged struct {
	_ byte // padding
}

// Experimental.
func NewFirewallAttachmentStatusChanged() FirewallAttachmentStatusChanged {
	_init_.Initialize()

	j := jsiiProxy_FirewallAttachmentStatusChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallAttachmentStatusChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallAttachmentStatusChanged_Override(f FirewallAttachmentStatusChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallAttachmentStatusChanged",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for Firewall Attachment Status Changed.
// Experimental.
func FirewallAttachmentStatusChanged_FirewallAttachmentStatusChangedPattern(options *FirewallAttachmentStatusChanged_FirewallAttachmentStatusChangedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFirewallAttachmentStatusChanged_FirewallAttachmentStatusChangedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallAttachmentStatusChanged",
		"firewallAttachmentStatusChangedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

