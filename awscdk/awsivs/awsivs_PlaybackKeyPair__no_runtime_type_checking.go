//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsivs

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PlaybackKeyPair) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PlaybackKeyPair) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PlaybackKeyPair) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PlaybackKeyPair) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_PlaybackKeyPair) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validatePlaybackKeyPair_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePlaybackKeyPair_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewPlaybackKeyPairParameters(scope constructs.Construct, id *string, props *PlaybackKeyPairProps) error {
	return nil
}

