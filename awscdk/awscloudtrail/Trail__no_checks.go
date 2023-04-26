//go:build no_runtime_type_checking

package awscloudtrail

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Trail) validateAddEventSelectorParameters(dataResourceType DataResourceType, dataResourceValues *[]*string, options *AddEventSelectorOptions) error {
	return nil
}

func (t *jsiiProxy_Trail) validateAddLambdaEventSelectorParameters(handlers *[]awslambda.IFunction, options *AddEventSelectorOptions) error {
	return nil
}

func (t *jsiiProxy_Trail) validateAddS3EventSelectorParameters(s3Selector *[]*S3EventSelector, options *AddEventSelectorOptions) error {
	return nil
}

func (t *jsiiProxy_Trail) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Trail) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Trail) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_Trail) validateLogAllLambdaDataEventsParameters(options *AddEventSelectorOptions) error {
	return nil
}

func (t *jsiiProxy_Trail) validateLogAllS3DataEventsParameters(options *AddEventSelectorOptions) error {
	return nil
}

func validateTrail_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTrail_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTrail_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTrail_OnEventParameters(scope constructs.Construct, id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateNewTrailParameters(scope constructs.Construct, id *string, props *TrailProps) error {
	return nil
}

