//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PublishAssetsAction) validateAddPublishCommandParameters(relativeManifestPath *string, assetSelector *string) error {
	return nil
}

func (p *jsiiProxy_PublishAssetsAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (p *jsiiProxy_PublishAssetsAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (p *jsiiProxy_PublishAssetsAction) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_PublishAssetsAction) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validatePublishAssetsAction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPublishAssetsActionParameters(scope constructs.Construct, id *string, props *PublishAssetsActionProps) error {
	return nil
}

