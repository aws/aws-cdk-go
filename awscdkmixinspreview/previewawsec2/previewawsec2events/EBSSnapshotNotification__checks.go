//go:build !no_runtime_type_checking

package previewawsec2events

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEBSSnapshotNotification_EbsSnapshotNotificationPatternParameters(options *EBSSnapshotNotification_EBSSnapshotNotificationProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

