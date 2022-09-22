//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsiot

import (
	"fmt"
)

func (i *jsiiProxy_IAction) validateBindParameters(topicRule ITopicRule) error {
	if topicRule == nil {
		return fmt.Errorf("parameter topicRule is required, but nil was provided")
	}

	return nil
}

