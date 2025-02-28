//go:build !no_runtime_type_checking

package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

func (a *jsiiProxy_AssetManifestBuilder) validateAddDockerImageAssetParameters(stack Stack, sourceHash *string, source *cloudassemblyschema.DockerImageSource, dest *cloudassemblyschema.DockerImageDestination) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if sourceHash == nil {
		return fmt.Errorf("parameter sourceHash is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(source, func() string { return "parameter source" }); err != nil {
		return err
	}

	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(dest, func() string { return "parameter dest" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateAddFileAssetParameters(stack Stack, sourceHash *string, source *cloudassemblyschema.FileSource, dest *cloudassemblyschema.FileDestination) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if sourceHash == nil {
		return fmt.Errorf("parameter sourceHash is required, but nil was provided")
	}

	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(source, func() string { return "parameter source" }); err != nil {
		return err
	}

	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(dest, func() string { return "parameter dest" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateDefaultAddDockerImageAssetParameters(stack Stack, asset *DockerImageAssetSource, target *AssetManifestDockerImageDestination) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(target, func() string { return "parameter target" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateDefaultAddFileAssetParameters(stack Stack, asset *FileAssetSource, target *AssetManifestFileDestination) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(target, func() string { return "parameter target" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AssetManifestBuilder) validateEmitManifestParameters(stack Stack, session ISynthesisSession, options *cloudassemblyschema.AssetManifestOptions) error {
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

