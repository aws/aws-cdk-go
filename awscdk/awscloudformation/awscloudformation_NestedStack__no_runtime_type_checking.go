//go:build no_runtime_type_checking

package awscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NestedStack) validateAddDependencyParameters(target awscdk.Stack) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAddDockerImageAssetParameters(asset *awscdk.DockerImageAssetSource) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAddFileAssetParameters(asset *awscdk.FileAssetSource) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateParseArnParameters(arn *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validatePrepareCrossReferenceParameters(_sourceStack awscdk.Stack, reference awscdk.Reference) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateReportMissingContextParameters(report *cxapi.MissingContext) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateResolveParameters(obj interface{}) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateSetParameterParameters(name *string, value *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateToJsonStringParameters(obj interface{}) error {
	return nil
}

func validateNestedStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNestedStack_IsNestedStackParameters(x interface{}) error {
	return nil
}

func validateNestedStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateNestedStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNestedStackParameters(scope awscdk.Construct, id *string, props *NestedStackProps) error {
	return nil
}

