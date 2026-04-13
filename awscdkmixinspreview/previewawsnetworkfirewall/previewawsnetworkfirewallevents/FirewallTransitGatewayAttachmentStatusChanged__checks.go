//go:build !no_runtime_type_checking

package previewawsnetworkfirewallevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateFirewallTransitGatewayAttachmentStatusChanged_EventPatternParameters(options *FirewallTransitGatewayAttachmentStatusChanged_FirewallTransitGatewayAttachmentStatusChangedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

