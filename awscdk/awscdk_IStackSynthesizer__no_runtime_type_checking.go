//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
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

