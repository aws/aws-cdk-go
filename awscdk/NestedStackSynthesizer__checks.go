//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

func (n *jsiiProxy_NestedStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	if requiredVersion == nil {
		return fmt.Errorf("parameter requiredVersion is required, but nil was provided")
	}

	if bootstrapStackVersionSsmParameter == nil {
		return fmt.Errorf("parameter bootstrapStackVersionSsmParameter is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateBindParameters(stack Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(dest, func() string { return "parameter dest" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(location, func() string { return "parameter location" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NestedStackSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateNewNestedStackSynthesizerParameters(parentDeployment IStackSynthesizer) error {
	if parentDeployment == nil {
		return fmt.Errorf("parameter parentDeployment is required, but nil was provided")
	}

	return nil
}

