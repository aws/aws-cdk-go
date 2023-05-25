//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Source) validateBindParameters(_scope constructs.Construct, _project IProject) error {
	return nil
}

func validateSource_BitBucketParameters(props *BitBucketSourceProps) error {
	return nil
}

func validateSource_CodeCommitParameters(props *CodeCommitSourceProps) error {
	return nil
}

func validateSource_GitHubParameters(props *GitHubSourceProps) error {
	return nil
}

func validateSource_GitHubEnterpriseParameters(props *GitHubEnterpriseSourceProps) error {
	return nil
}

func validateSource_S3Parameters(props *S3SourceProps) error {
	return nil
}

func validateNewSourceParameters(props *SourceProps) error {
	return nil
}

