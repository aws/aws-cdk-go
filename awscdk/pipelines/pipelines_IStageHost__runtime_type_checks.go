//go:build !no_runtime_type_checking

package pipelines

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IStageHost) validatePublishAssetParameters(command *AssetPublishingCommand) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(command, func() string { return "parameter command" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IStageHost) validateStackOutputArtifactParameters(stackArtifactId *string) error {
	if stackArtifactId == nil {
		return fmt.Errorf("parameter stackArtifactId is required, but nil was provided")
	}

	return nil
}

