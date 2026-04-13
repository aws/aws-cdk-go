//go:build !no_runtime_type_checking

package previewawsroute53resolverevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDNSFirewallBlock_EventPatternParameters(options *DNSFirewallBlock_DNSFirewallBlockProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

