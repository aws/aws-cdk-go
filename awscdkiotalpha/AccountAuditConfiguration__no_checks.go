//go:build no_runtime_type_checking

package awscdkiotalpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountAuditConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AccountAuditConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AccountAuditConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAccountAuditConfiguration_FromAccountIdParameters(scope constructs.Construct, id *string, accountId *string) error {
	return nil
}

func validateAccountAuditConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAccountAuditConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAccountAuditConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAccountAuditConfigurationParameters(scope constructs.Construct, id *string, props *AccountAuditConfigurationProps) error {
	return nil
}

