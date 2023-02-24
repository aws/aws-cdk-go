//go:build !no_runtime_type_checking

package assertions

import (
	"fmt"
)

func (m *jsiiProxy_Matcher) validateTestParameters(actual interface{}) error {
	if actual == nil {
		return fmt.Errorf("parameter actual is required, but nil was provided")
	}

	return nil
}

func validateMatcher_IsMatcherParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

