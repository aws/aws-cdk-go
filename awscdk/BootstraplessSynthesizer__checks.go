//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

func (b *jsiiProxy_BootstraplessSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	if requiredVersion == nil {
		return fmt.Errorf("parameter requiredVersion is required, but nil was provided")
	}

	if bootstrapStackVersionSsmParameter == nil {
		return fmt.Errorf("parameter bootstrapStackVersionSsmParameter is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateAddDockerImageAssetParameters(_asset *DockerImageAssetSource) error {
	if _asset == nil {
		return fmt.Errorf("parameter _asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_asset, func() string { return "parameter _asset" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateAddFileAssetParameters(_asset *FileAssetSource) error {
	if _asset == nil {
		return fmt.Errorf("parameter _asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_asset, func() string { return "parameter _asset" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateBindParameters(stack Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(dest, func() string { return "parameter dest" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(location, func() string { return "parameter location" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
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

func (b *jsiiProxy_BootstraplessSynthesizer) validateReusableBindParameters(stack Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BootstraplessSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateNewBootstraplessSynthesizerParameters(props *BootstraplessSynthesizerProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

