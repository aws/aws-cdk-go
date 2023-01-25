//go:build !no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

func validateNewBootstraplessSynthesizerParameters(props *BootstraplessSynthesizerProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

