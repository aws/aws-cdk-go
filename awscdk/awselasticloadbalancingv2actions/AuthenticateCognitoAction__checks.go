//go:build !no_runtime_type_checking

package awselasticloadbalancingv2actions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AuthenticateCognitoAction) validateAddRuleActionParameters(actionJson *awselasticloadbalancingv2.CfnListenerRule_ActionProperty) error {
	if actionJson == nil {
		return fmt.Errorf("parameter actionJson is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(actionJson, func() string { return "parameter actionJson" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AuthenticateCognitoAction) validateBindParameters(scope constructs.Construct, listener awselasticloadbalancingv2.IApplicationListener) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if listener == nil {
		return fmt.Errorf("parameter listener is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AuthenticateCognitoAction) validateRenumberParameters(actions *[]*awselasticloadbalancingv2.CfnListener_ActionProperty) error {
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

func validateAuthenticateCognitoAction_AuthenticateOidcParameters(options *awselasticloadbalancingv2.AuthenticateOidcOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAuthenticateCognitoAction_FixedResponseParameters(statusCode *float64, options *awselasticloadbalancingv2.FixedResponseOptions) error {
	if statusCode == nil {
		return fmt.Errorf("parameter statusCode is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAuthenticateCognitoAction_ForwardParameters(targetGroups *[]awselasticloadbalancingv2.IApplicationTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) error {
	if targetGroups == nil {
		return fmt.Errorf("parameter targetGroups is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAuthenticateCognitoAction_RedirectParameters(options *awselasticloadbalancingv2.RedirectOptions) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateAuthenticateCognitoAction_WeightedForwardParameters(targetGroups *[]*awselasticloadbalancingv2.WeightedTargetGroup, options *awselasticloadbalancingv2.ForwardOptions) error {
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

func validateNewAuthenticateCognitoActionParameters(options *AuthenticateCognitoActionProps) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

