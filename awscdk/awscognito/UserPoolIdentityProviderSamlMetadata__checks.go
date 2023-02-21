//go:build !no_runtime_type_checking

package awscognito

import (
	"fmt"
)

func validateUserPoolIdentityProviderSamlMetadata_FileParameters(fileContent *string) error {
	if fileContent == nil {
		return fmt.Errorf("parameter fileContent is required, but nil was provided")
	}

	return nil
}

func validateUserPoolIdentityProviderSamlMetadata_UrlParameters(url *string) error {
	if url == nil {
		return fmt.Errorf("parameter url is required, but nil was provided")
	}

	return nil
}

