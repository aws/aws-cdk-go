//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (a *jsiiProxy_ApplicationListenerRule) validateAddConditionParameters(condition ListenerCondition) error {
	if condition == nil {
		return fmt.Errorf("parameter condition is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateAddFixedResponseParameters(fixedResponse *FixedResponse) error {
	if fixedResponse == nil {
		return fmt.Errorf("parameter fixedResponse is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(fixedResponse, func() string { return "parameter fixedResponse" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateAddRedirectResponseParameters(redirectResponse *RedirectResponse) error {
	if redirectResponse == nil {
		return fmt.Errorf("parameter redirectResponse is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(redirectResponse, func() string { return "parameter redirectResponse" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateAddTargetGroupParameters(targetGroup IApplicationTargetGroup) error {
	if targetGroup == nil {
		return fmt.Errorf("parameter targetGroup is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateConfigureActionParameters(action ListenerAction) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateSetConditionParameters(field *string) error {
	if field == nil {
		return fmt.Errorf("parameter field is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApplicationListenerRule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateApplicationListenerRule_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewApplicationListenerRuleParameters(scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

