//go:build !no_runtime_type_checking

package cxapi

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

func (c *jsiiProxy_CloudAssemblyBuilder) validateAddArtifactParameters(id *string, manifest *cloudassemblyschema.ArtifactManifest) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if manifest == nil {
		return fmt.Errorf("parameter manifest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(manifest, func() string { return "parameter manifest" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudAssemblyBuilder) validateAddMissingParameters(missing *cloudassemblyschema.MissingContext) error {
	if missing == nil {
		return fmt.Errorf("parameter missing is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(missing, func() string { return "parameter missing" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudAssemblyBuilder) validateBuildAssemblyParameters(options *AssemblyBuildOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CloudAssemblyBuilder) validateCreateNestedAssemblyParameters(artifactId *string, displayName *string) error {
	if artifactId == nil {
		return fmt.Errorf("parameter artifactId is required, but nil was provided")
	}

	if displayName == nil {
		return fmt.Errorf("parameter displayName is required, but nil was provided")
	}

	return nil
}

func validateNewCloudAssemblyBuilderParameters(props *CloudAssemblyBuilderProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

