//go:build !no_runtime_type_checking

package integtests

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (a *jsiiProxy_AssertionsProvider) validateAddPolicyStatementFromSdkCallParameters(service *string, api *string) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateEncodeParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AssertionsProvider) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateAssertionsProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAssertionsProviderParameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

