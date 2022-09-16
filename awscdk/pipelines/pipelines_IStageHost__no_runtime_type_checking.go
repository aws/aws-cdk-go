//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStageHost) validatePublishAssetParameters(command *AssetPublishingCommand) error {
	return nil
}

func (i *jsiiProxy_IStageHost) validateStackOutputArtifactParameters(stackArtifactId *string) error {
	return nil
}

