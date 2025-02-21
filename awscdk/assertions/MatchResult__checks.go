//go:build !no_runtime_type_checking

package assertions

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_MatchResult) validateComposeParameters(id *string, inner MatchResult) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if inner == nil {
		return fmt.Errorf("parameter inner is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MatchResult) validatePushParameters(matcher Matcher, path *[]*string, message *string) error {
	if matcher == nil {
		return fmt.Errorf("parameter matcher is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if message == nil {
		return fmt.Errorf("parameter message is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MatchResult) validateRecordCaptureParameters(options *MatchCapture) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MatchResult) validateRecordFailureParameters(failure *MatchFailure) error {
	if failure == nil {
		return fmt.Errorf("parameter failure is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(failure, func() string { return "parameter failure" }); err != nil {
		return err
	}

	return nil
}

func validateNewMatchResultParameters(target interface{}) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

