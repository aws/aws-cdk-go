//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SourceApiAssociation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SourceApiAssociation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SourceApiAssociation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSourceApiAssociation_FromSourceApiAssociationAttributesParameters(scope constructs.Construct, id *string, attrs *SourceApiAssociationAttributes) error {
	return nil
}

func validateSourceApiAssociation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSourceApiAssociation_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSourceApiAssociation_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSourceApiAssociationParameters(scope constructs.Construct, id *string, props *SourceApiAssociationProps) error {
	return nil
}

