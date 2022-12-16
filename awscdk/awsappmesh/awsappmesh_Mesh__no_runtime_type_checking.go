//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Mesh) validateAddVirtualGatewayParameters(id *string, props *VirtualGatewayBaseProps) error {
	return nil
}

func (m *jsiiProxy_Mesh) validateAddVirtualNodeParameters(id *string, props *VirtualNodeBaseProps) error {
	return nil
}

func (m *jsiiProxy_Mesh) validateAddVirtualRouterParameters(id *string, props *VirtualRouterBaseProps) error {
	return nil
}

func (m *jsiiProxy_Mesh) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_Mesh) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_Mesh) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateMesh_FromMeshArnParameters(scope constructs.Construct, id *string, meshArn *string) error {
	return nil
}

func validateMesh_FromMeshNameParameters(scope constructs.Construct, id *string, meshName *string) error {
	return nil
}

func validateMesh_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMesh_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateMesh_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewMeshParameters(scope constructs.Construct, id *string, props *MeshProps) error {
	return nil
}

