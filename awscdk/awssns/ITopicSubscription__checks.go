//go:build !no_runtime_type_checking

package awssns

import (
	"fmt"
)

func (i *jsiiProxy_ITopicSubscription) validateBindParameters(topic ITopic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

