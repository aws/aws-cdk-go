//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualNode) validateAddBackendParameters(backend Backend) error {
	return nil
}

func (v *jsiiProxy_VirtualNode) validateAddListenerParameters(listener VirtualNodeListener) error {
	return nil
}

func (v *jsiiProxy_VirtualNode) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VirtualNode) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VirtualNode) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VirtualNode) validateGrantStreamAggregatedResourcesParameters(identity awsiam.IGrantable) error {
	return nil
}

func validateVirtualNode_FromVirtualNodeArnParameters(scope constructs.Construct, id *string, virtualNodeArn *string) error {
	return nil
}

func validateVirtualNode_FromVirtualNodeAttributesParameters(scope constructs.Construct, id *string, attrs *VirtualNodeAttributes) error {
	return nil
}

func validateVirtualNode_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVirtualNode_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVirtualNode_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVirtualNodeParameters(scope constructs.Construct, id *string, props *VirtualNodeProps) error {
	return nil
}

