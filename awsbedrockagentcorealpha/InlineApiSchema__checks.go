//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_InlineApiSchema) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InlineApiSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func validateInlineApiSchema_FromInlineParameters(schema *string) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

func validateInlineApiSchema_FromLocalAssetParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateInlineApiSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	return nil
}

func validateNewInlineApiSchemaParameters(schema *string) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

