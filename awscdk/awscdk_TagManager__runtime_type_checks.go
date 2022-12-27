//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (t *jsiiProxy_TagManager) validateRemoveTagParameters(key *string, priority *float64) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if priority == nil {
		return fmt.Errorf("parameter priority is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TagManager) validateSetTagParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTagManager_IsTaggableParameters(construct interface{}) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewTagManagerParameters(tagType TagType, resourceTypeName *string, options *TagManagerOptions) error {
	if tagType == "" {
		return fmt.Errorf("parameter tagType is required, but nil was provided")
	}

	if resourceTypeName == nil {
		return fmt.Errorf("parameter resourceTypeName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

