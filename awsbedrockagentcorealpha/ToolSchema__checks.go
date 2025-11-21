//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_ToolSchema) validateBindParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_ToolSchema) validateGrantPermissionsToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func validateToolSchema_FromInlineParameters(schema *[]*ToolDefinition) error {
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

func validateToolSchema_FromLocalAssetParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateToolSchema_FromS3FileParameters(bucket awss3.IBucket, objectKey *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if objectKey == nil {
		return fmt.Errorf("parameter objectKey is required, but nil was provided")
	}

	return nil
}

func validateNewToolSchemaParameters(s3File *awss3.Location, inlineSchema *[]*ToolDefinition) error {
	if err := _jsii_.ValidateStruct(s3File, func() string { return "parameter s3File" }); err != nil {
		return err
	}

	for idx_6268cf, v := range *inlineSchema {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter inlineSchema[%#v]", idx_6268cf) }); err != nil {
			return err
		}
	}

	return nil
}

