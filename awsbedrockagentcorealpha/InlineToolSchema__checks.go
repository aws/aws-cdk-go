//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_InlineToolSchema) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_InlineToolSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func validateInlineToolSchema_FromInlineParameters(schema *[]*ToolDefinition) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}
	for idx_df0ad6, v := range *schema {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter schema[%#v]", idx_df0ad6) }); err != nil {
			return err
		}
	}

	return nil
}

func validateInlineToolSchema_FromLocalAssetParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateInlineToolSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	return nil
}

func validateNewInlineToolSchemaParameters(schema *[]*ToolDefinition) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}
	for idx_df0ad6, v := range *schema {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter schema[%#v]", idx_df0ad6) }); err != nil {
			return err
		}
	}

	return nil
}

