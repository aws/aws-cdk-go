//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventSourceMapping) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EventSourceMapping) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EventSourceMapping) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateEventSourceMapping_FromEventSourceMappingIdParameters(scope constructs.Construct, id *string, eventSourceMappingId *string) error {
	return nil
}

func validateEventSourceMapping_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventSourceMapping_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEventSourceMapping_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEventSourceMappingParameters(scope constructs.Construct, id *string, props *EventSourceMappingProps) error {
	return nil
}

