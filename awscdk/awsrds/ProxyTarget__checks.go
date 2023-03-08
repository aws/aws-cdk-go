//go:build !no_runtime_type_checking

package awsrds

import (
	"fmt"
)

func (p *jsiiProxy_ProxyTarget) validateBindParameters(proxy DatabaseProxy) error {
	if proxy == nil {
		return fmt.Errorf("parameter proxy is required, but nil was provided")
	}

	return nil
}

func validateProxyTarget_FromClusterParameters(cluster IDatabaseCluster) error {
	if cluster == nil {
		return fmt.Errorf("parameter cluster is required, but nil was provided")
	}

	return nil
}

func validateProxyTarget_FromInstanceParameters(instance IDatabaseInstance) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

