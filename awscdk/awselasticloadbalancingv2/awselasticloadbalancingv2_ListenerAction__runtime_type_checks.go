//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (l *jsiiProxy_ListenerAction) validateBindParameters(scope constructs.Construct, listener IApplicationListener) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_ListenerAction) validateRenumberParameters(actions *[]*CfnListener_ActionProperty) error {
	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}
	for idx_2b0dcd, v := range *actions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter actions[%#v]", idx_2b0dcd) }); err != nil {
			return err
		}
	}

	return nil
}

func validateListenerAction_AuthenticateOidcParameters(options *AuthenticateOidcOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateListenerAction_FixedResponseParameters(statusCode *float64, options *FixedResponseOptions) error {
	if statusCode == nil {
		return fmt.Errorf("parameter statusCode is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateListenerAction_ForwardParameters(targetGroups *[]IApplicationTargetGroup, options *ForwardOptions) error {
	if targetGroups == nil {
		return fmt.Errorf("parameter targetGroups is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateListenerAction_RedirectParameters(options *RedirectOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateListenerAction_WeightedForwardParameters(targetGroups *[]*WeightedTargetGroup, options *ForwardOptions) error {
	if targetGroups == nil {
		return fmt.Errorf("parameter targetGroups is required, but nil was provided")
	}
	for idx_48834c, v := range *targetGroups {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter targetGroups[%#v]", idx_48834c) }); err != nil {
			return err
		}
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewListenerActionParameters(actionJson *CfnListener_ActionProperty) error {
	if actionJson == nil {
		return fmt.Errorf("parameter actionJson is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(actionJson, func() string { return "parameter actionJson" }); err != nil {
		return err
	}

	return nil
}

