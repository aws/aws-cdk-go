//go:build !no_runtime_type_checking

package awsecrmixins

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (r *jsiiProxy_RepositoryAutoDeleteImages) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RepositoryAutoDeleteImages) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateRepositoryAutoDeleteImages_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

