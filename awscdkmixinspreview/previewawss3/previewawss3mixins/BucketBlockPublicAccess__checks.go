//go:build !no_runtime_type_checking

package previewawss3mixins

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BucketBlockPublicAccess) validateApplyToParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBlockPublicAccess) validateSupportsParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

