//go:build !no_runtime_type_checking

package awsiam

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (p *jsiiProxy_PolicyStatement) validateAddAccountConditionParameters(accountId *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddArnPrincipalParameters(arn *string) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddAwsAccountPrincipalParameters(accountId *string) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddCanonicalUserPrincipalParameters(canonicalUserId *string) error {
	if canonicalUserId == nil {
		return fmt.Errorf("parameter canonicalUserId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddConditionParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddConditionsParameters(conditions *map[string]interface{}) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddFederatedPrincipalParameters(federated interface{}, conditions *map[string]interface{}) error {
	if federated == nil {
		return fmt.Errorf("parameter federated is required, but nil was provided")
	}

	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateAddServicePrincipalParameters(service *string, opts *ServicePrincipalOpts) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateCopyParameters(overrides *PolicyStatementProps) error {
	if err := _jsii_.ValidateStruct(overrides, func() string { return "parameter overrides" }); err != nil {
		return err
	}

	return nil
}

func validatePolicyStatement_FromJsonParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PolicyStatement) validateSetEffectParameters(val Effect) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPolicyStatementParameters(props *PolicyStatementProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

