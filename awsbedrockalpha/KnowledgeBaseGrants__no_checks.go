//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KnowledgeBaseGrants) validateActionsParameters(grantee awsiam.IGrantable, actions *[]*string, options *awscdk.PermissionsOptions) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseGrants) validateGetDocumentContentParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseGrants) validateRetrieveParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBaseGrants) validateRetrieveAndGenerateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateKnowledgeBaseGrants_FromKnowledgeBaseParameters(resource interfacesawsbedrock.IKnowledgeBaseRef) error {
	return nil
}

