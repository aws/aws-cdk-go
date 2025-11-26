//go:build !no_runtime_type_checking

package previewawsconnectevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect"
)

func (i *jsiiProxy_InstanceEvents) validateCodeConnectContactPatternParameters(options *InstanceEvents_CodeConnectContact_CodeConnectContactProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_InstanceEvents) validateContactLensPostCallRulesMatchedPatternParameters(options *InstanceEvents_ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_InstanceEvents) validateContactLensRealtimeRulesMatchedPatternParameters(options *InstanceEvents_ContactLensRealtimeRulesMatched_ContactLensRealtimeRulesMatchedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateInstanceEvents_FromInstanceParameters(instanceRef interfacesawsconnect.IInstanceRef) error {
	if instanceRef == nil {
		return fmt.Errorf("parameter instanceRef is required, but nil was provided")
	}

	return nil
}

