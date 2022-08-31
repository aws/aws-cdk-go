//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	return nil
}

func (i *jsiiProxy_IStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	return nil
}

func (i *jsiiProxy_IStackSynthesizer) validateBindParameters(stack Stack) error {
	return nil
}

func (i *jsiiProxy_IStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	return nil
}

