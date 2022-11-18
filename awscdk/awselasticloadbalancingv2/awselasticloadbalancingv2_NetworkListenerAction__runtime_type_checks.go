//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NetworkListenerAction) validateBindParameters(scope constructs.Construct, listener INetworkListener) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkListenerAction) validateRenumberParameters(actions *[]*CfnListener_ActionProperty) error {
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

func validateNetworkListenerAction_ForwardParameters(targetGroups *[]INetworkTargetGroup, options *NetworkForwardOptions) error {
	if targetGroups == nil {
		return fmt.Errorf("parameter targetGroups is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNetworkListenerAction_WeightedForwardParameters(targetGroups *[]*NetworkWeightedTargetGroup, options *NetworkForwardOptions) error {
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

func validateNewNetworkListenerActionParameters(actionJson *CfnListener_ActionProperty) error {
	if actionJson == nil {
		return fmt.Errorf("parameter actionJson is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(actionJson, func() string { return "parameter actionJson" }); err != nil {
		return err
	}

	return nil
}

