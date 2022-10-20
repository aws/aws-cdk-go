//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stack) validateAddDependencyParameters(target Stack) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAllocateLogicalIdParameters(cfnElement CfnElement) error {
	return nil
}

func (s *jsiiProxy_Stack) validateExportValueParameters(exportedValue interface{}, options *ExportValueOptions) error {
	return nil
}

func (s *jsiiProxy_Stack) validateFormatArnParameters(components *ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetLogicalIdParameters(element CfnElement) error {
	return nil
}

func (s *jsiiProxy_Stack) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Stack) validateParseArnParameters(arn *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validatePrepareCrossReferenceParameters(_sourceStack Stack, reference Reference) error {
	return nil
}

func (s *jsiiProxy_Stack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateReportMissingContextParameters(report *cxapi.MissingContext) error {
	return nil
}

func (s *jsiiProxy_Stack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	return nil
}

func (s *jsiiProxy_Stack) validateResolveParameters(obj interface{}) error {
	return nil
}

func (s *jsiiProxy_Stack) validateSplitArnParameters(arn *string, arnFormat ArnFormat) error {
	return nil
}

func (s *jsiiProxy_Stack) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Stack) validateToJsonStringParameters(obj interface{}) error {
	return nil
}

func validateStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStack_IsStackParameters(x interface{}) error {
	return nil
}

func validateStack_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStackParameters(props *StackProps) error {
	return nil
}

