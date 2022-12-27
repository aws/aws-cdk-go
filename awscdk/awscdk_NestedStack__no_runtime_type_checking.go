//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NestedStack) validateAddDependencyParameters(target Stack) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAddMetadataParameters(key *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAddTransformParameters(transform *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateAllocateLogicalIdParameters(cfnElement CfnElement) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateExportStringListValueParameters(exportedValue interface{}, options *ExportValueOptions) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateExportValueParameters(exportedValue interface{}, options *ExportValueOptions) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateFormatArnParameters(components *ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateGetLogicalIdParameters(element CfnElement) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateRegionalFactParameters(factName *string) error {
	return nil
}

func (n *jsiiProxy_NestedStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
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

func (n *jsiiProxy_NestedStack) validateSplitArnParameters(arn *string, arnFormat ArnFormat) error {
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

func validateNewNestedStackParameters(scope constructs.Construct, id *string, props *NestedStackProps) error {
	return nil
}

