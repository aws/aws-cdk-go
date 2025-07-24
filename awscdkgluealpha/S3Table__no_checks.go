//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3Table) validateAddPartitionIndexParameters(index *PartitionIndex) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGrantParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGrantReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGrantToUnderlyingResourcesParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (s *jsiiProxy_S3Table) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateS3Table_FromTableArnParameters(scope constructs.Construct, id *string, tableArn *string) error {
	return nil
}

func validateS3Table_FromTableAttributesParameters(scope constructs.Construct, id *string, attrs *TableAttributes) error {
	return nil
}

func validateS3Table_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3Table_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateS3Table_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewS3TableParameters(scope constructs.Construct, id *string, props *S3TableProps) error {
	return nil
}

