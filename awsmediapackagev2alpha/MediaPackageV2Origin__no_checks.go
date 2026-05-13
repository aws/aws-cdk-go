//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MediaPackageV2Origin) validateBindParameters(scope constructs.Construct, options *awscloudfront.OriginBindOptions) error {
	return nil
}

func validateNewMediaPackageV2OriginParameters(endpoint IOriginEndpoint, props *MediaPackageV2OriginProps) error {
	return nil
}

