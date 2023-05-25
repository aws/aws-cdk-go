//go:build no_runtime_type_checking

package assertions

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MatchResult) validateComposeParameters(id *string, inner MatchResult) error {
	return nil
}

func (m *jsiiProxy_MatchResult) validatePushParameters(matcher Matcher, path *[]*string, message *string) error {
	return nil
}

func (m *jsiiProxy_MatchResult) validateRecordCaptureParameters(options *MatchCapture) error {
	return nil
}

func (m *jsiiProxy_MatchResult) validateRecordFailureParameters(failure *MatchFailure) error {
	return nil
}

func validateNewMatchResultParameters(target interface{}) error {
	return nil
}

