//go:build !no_runtime_type_checking

package previewawshealthimagingevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawshealthimaging"
)

func (d *jsiiProxy_DatastoreEvents) validateDataStoreCreatedPatternParameters(options *DataStoreCreated_DataStoreCreatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreCreatingPatternParameters(options *DataStoreCreating_DataStoreCreatingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreCreationFailedPatternParameters(options *DataStoreCreationFailed_DataStoreCreationFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreDeletedPatternParameters(options *DataStoreDeleted_DataStoreDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreDeletingPatternParameters(options *DataStoreDeleting_DataStoreDeletingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopiedPatternParameters(options *ImageSetCopied_ImageSetCopiedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopyFailedPatternParameters(options *ImageSetCopyFailed_ImageSetCopyFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopyingPatternParameters(options *ImageSetCopying_ImageSetCopyingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopyingWithReadOnlyAccessPatternParameters(options *ImageSetCopyingWithReadOnlyAccess_ImageSetCopyingWithReadOnlyAccessProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCreatedPatternParameters(options *ImageSetCreated_ImageSetCreatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetDeletedPatternParameters(options *ImageSetDeleted_ImageSetDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetDeletingPatternParameters(options *ImageSetDeleting_ImageSetDeletingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetUpdatedPatternParameters(options *ImageSetUpdated_ImageSetUpdatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetUpdateFailedPatternParameters(options *ImageSetUpdateFailed_ImageSetUpdateFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetUpdatingPatternParameters(options *ImageSetUpdating_ImageSetUpdatingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobCompletedPatternParameters(options *ImportJobCompleted_ImportJobCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobFailedPatternParameters(options *ImportJobFailed_ImportJobFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobInProgressPatternParameters(options *ImportJobInProgress_ImportJobInProgressProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobSubmittedPatternParameters(options *ImportJobSubmitted_ImportJobSubmittedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDatastoreEvents_FromDatastoreParameters(datastoreRef interfacesawshealthimaging.IDatastoreRef) error {
	if datastoreRef == nil {
		return fmt.Errorf("parameter datastoreRef is required, but nil was provided")
	}

	return nil
}

