//go:build !no_runtime_type_checking

package previewawsnetworkmanagerevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager"
)

func (c *jsiiProxy_CoreNetworkEvents) validateNetworkManagerPolicyUpdatePatternParameters(options *CoreNetworkEvents_NetworkManagerPolicyUpdate_NetworkManagerPolicyUpdateProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CoreNetworkEvents) validateNetworkManagerSegmentUpdatePatternParameters(options *CoreNetworkEvents_NetworkManagerSegmentUpdate_NetworkManagerSegmentUpdateProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateCoreNetworkEvents_FromCoreNetworkParameters(coreNetworkRef interfacesawsnetworkmanager.ICoreNetworkRef) error {
	if coreNetworkRef == nil {
		return fmt.Errorf("parameter coreNetworkRef is required, but nil was provided")
	}

	return nil
}

