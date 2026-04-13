package previewawsnetworkfirewallevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.networkfirewall@FirewallConfigurationChanged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   firewallConfigurationChanged := awscdkmixinspreview.Events.NewFirewallConfigurationChanged()
//
// Experimental.
type FirewallConfigurationChanged interface {
}

// The jsii proxy struct for FirewallConfigurationChanged
type jsiiProxy_FirewallConfigurationChanged struct {
	_ byte // padding
}

// Experimental.
func NewFirewallConfigurationChanged() FirewallConfigurationChanged {
	_init_.Initialize()

	j := jsiiProxy_FirewallConfigurationChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallConfigurationChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFirewallConfigurationChanged_Override(f FirewallConfigurationChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallConfigurationChanged",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for Firewall Configuration Changed.
// Experimental.
func FirewallConfigurationChanged_EventPattern(options *FirewallConfigurationChanged_FirewallConfigurationChangedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFirewallConfigurationChanged_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.events.FirewallConfigurationChanged",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

