//go:build !no_runtime_type_checking

package previewawsroute53resolverevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDNSFirewallAlert_EventPatternParameters(options *DNSFirewallAlert_DNSFirewallAlertProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

