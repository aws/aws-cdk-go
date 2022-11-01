//go:build !no_runtime_type_checking

package cxapi

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

func (t *jsiiProxy_TreeCloudArtifact) validateFindMetadataByTypeParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validateTreeCloudArtifact_FromManifestParameters(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	if assembly == nil {
		return fmt.Errorf("parameter assembly is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if artifact == nil {
		return fmt.Errorf("parameter artifact is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(artifact, func() string { return "parameter artifact" }); err != nil {
		return err
	}

	return nil
}

func validateNewTreeCloudArtifactParameters(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) error {
	if assembly == nil {
		return fmt.Errorf("parameter assembly is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if artifact == nil {
		return fmt.Errorf("parameter artifact is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(artifact, func() string { return "parameter artifact" }); err != nil {
		return err
	}

	return nil
}

