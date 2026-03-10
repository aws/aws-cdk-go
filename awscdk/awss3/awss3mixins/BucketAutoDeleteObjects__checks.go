//go:build !no_runtime_type_checking

package awss3mixins

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BucketAutoDeleteObjects) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketAutoDeleteObjects) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBucketAutoDeleteObjects_IsMixinParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

