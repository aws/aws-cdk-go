//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobDefinition) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (j *jsiiProxy_JobDefinition) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (j *jsiiProxy_JobDefinition) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateJobDefinition_FromJobDefinitionArnParameters(scope constructs.Construct, id *string, jobDefinitionArn *string) error {
	return nil
}

func validateJobDefinition_FromJobDefinitionNameParameters(scope constructs.Construct, id *string, jobDefinitionName *string) error {
	return nil
}

func validateJobDefinition_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJobDefinition_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateJobDefinition_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewJobDefinitionParameters(scope constructs.Construct, id *string, props *JobDefinitionProps) error {
	return nil
}

