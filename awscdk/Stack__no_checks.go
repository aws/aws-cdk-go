//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stack) validateAddDependencyParameters(target Stack) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAllocateLogicalIdParameters(cfnElement CfnElement) error {
	return nil
}

func (s *jsiiProxy_Stack) validateExportStringListValueParameters(exportedValue interface{}, options *ExportValueOptions) error {
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

func (s *jsiiProxy_Stack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
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

