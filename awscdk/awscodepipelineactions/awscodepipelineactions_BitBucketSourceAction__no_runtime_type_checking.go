//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BitBucketSourceAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (b *jsiiProxy_BitBucketSourceAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func validateNewBitBucketSourceActionParameters(props *BitBucketSourceActionProps) error {
	return nil
}

