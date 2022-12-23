//go:build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegTestCaseStack) validateAddDependencyParameters(target awscdk.Stack) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateAddDockerImageAssetParameters(asset *awscdk.DockerImageAssetSource) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateAddFileAssetParameters(asset *awscdk.FileAssetSource) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateParseArnParameters(arn *string) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validatePrepareCrossReferenceParameters(_sourceStack awscdk.Stack, reference awscdk.Reference) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateReportMissingContextParameters(report *cxapi.MissingContext) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateResolveParameters(obj interface{}) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_IntegTestCaseStack) validateToJsonStringParameters(obj interface{}) error {
	return nil
}

func validateIntegTestCaseStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIntegTestCaseStack_IsIntegTestCaseStackParameters(x interface{}) error {
	return nil
}

func validateIntegTestCaseStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateIntegTestCaseStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewIntegTestCaseStackParameters(scope constructs.Construct, id *string, props *IntegTestCaseStackProps) error {
	return nil
}

