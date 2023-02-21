//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBus) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (e *jsiiProxy_EventBus) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EventBus) validateArchiveParameters(id *string, props *BaseArchiveProps) error {
	return nil
}

func (e *jsiiProxy_EventBus) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EventBus) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EventBus) validateGrantPutEventsToParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEventBus_FromEventBusArnParameters(scope constructs.Construct, id *string, eventBusArn *string) error {
	return nil
}

func validateEventBus_FromEventBusAttributesParameters(scope constructs.Construct, id *string, attrs *EventBusAttributes) error {
	return nil
}

func validateEventBus_FromEventBusNameParameters(scope constructs.Construct, id *string, eventBusName *string) error {
	return nil
}

func validateEventBus_GrantAllPutEventsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEventBus_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventBus_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEventBus_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEventBusParameters(scope constructs.Construct, id *string, props *EventBusProps) error {
	return nil
}

