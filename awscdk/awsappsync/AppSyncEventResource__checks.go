//go:build !no_runtime_type_checking

package awsappsync

import (
	"fmt"
)

func (a *jsiiProxy_AppSyncEventResource) validateResourceArnsParameters(api EventApiBase) error {
	if api == nil {
		return fmt.Errorf("parameter api is required, but nil was provided")
	}

	return nil
}

func validateAppSyncEventResource_OfChannelNamespaceParameters(channelNamespace *string) error {
	if channelNamespace == nil {
		return fmt.Errorf("parameter channelNamespace is required, but nil was provided")
	}

	return nil
}

