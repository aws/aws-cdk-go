//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func (p *jsiiProxy_PolicyStatement) validateForPrincipalParameters(entityType *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateForPrincipalInGroupParameters(groupType *string, groupId *string) error {
	if groupType == nil {
		return fmt.Errorf("parameter groupType is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnActionParameters(action *string) error {
	if action == nil {
		return fmt.Errorf("parameter action is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnActionsParameters(actions *[]*string) error {
	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnResourceParameters(entityType *string, entityArn *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	if entityArn == nil {
		return fmt.Errorf("parameter entityArn is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PolicyStatement) validateOnResourceTypeParameters(entityType *string) error {
	if entityType == nil {
		return fmt.Errorf("parameter entityType is required, but nil was provided")
	}

	return nil
}

func validatePolicyStatement_FromCedarParameters(cedarStatement *string) error {
	if cedarStatement == nil {
		return fmt.Errorf("parameter cedarStatement is required, but nil was provided")
	}

	return nil
}

