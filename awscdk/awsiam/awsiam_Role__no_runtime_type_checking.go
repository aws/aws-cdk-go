//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Role) validateAddManagedPolicyParameters(policy IManagedPolicy) error {
	return nil
}

func (r *jsiiProxy_Role) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_Role) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_Role) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Role) validateAttachInlinePolicyParameters(policy Policy) error {
	return nil
}

func (r *jsiiProxy_Role) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Role) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_Role) validateGrantParameters(grantee IPrincipal) error {
	return nil
}

func (r *jsiiProxy_Role) validateGrantAssumeRoleParameters(identity IPrincipal) error {
	return nil
}

func (r *jsiiProxy_Role) validateGrantPassRoleParameters(identity IPrincipal) error {
	return nil
}

func (r *jsiiProxy_Role) validateWithoutPolicyUpdatesParameters(options *WithoutPolicyUpdatesOptions) error {
	return nil
}

func validateRole_CustomizeRolesParameters(scope constructs.Construct, options *CustomizeRolesOptions) error {
	return nil
}

func validateRole_FromRoleArnParameters(scope constructs.Construct, id *string, roleArn *string, options *FromRoleArnOptions) error {
	return nil
}

func validateRole_FromRoleNameParameters(scope constructs.Construct, id *string, roleName *string, options *FromRoleNameOptions) error {
	return nil
}

func validateRole_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRole_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRole_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRoleParameters(scope constructs.Construct, id *string, props *RoleProps) error {
	return nil
}

