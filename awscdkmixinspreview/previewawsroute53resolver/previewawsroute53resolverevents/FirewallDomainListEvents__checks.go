//go:build !no_runtime_type_checking

package previewawsroute53resolverevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53resolver"
)

func (f *jsiiProxy_FirewallDomainListEvents) validateDNSFirewallAlertPatternParameters(options *FirewallDomainListEvents_DNSFirewallAlert_DNSFirewallAlertProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (f *jsiiProxy_FirewallDomainListEvents) validateDNSFirewallBlockPatternParameters(options *FirewallDomainListEvents_DNSFirewallBlock_DNSFirewallBlockProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateFirewallDomainListEvents_FromFirewallDomainListParameters(firewallDomainListRef interfacesawsroute53resolver.IFirewallDomainListRef) error {
	if firewallDomainListRef == nil {
		return fmt.Errorf("parameter firewallDomainListRef is required, but nil was provided")
	}

	return nil
}

