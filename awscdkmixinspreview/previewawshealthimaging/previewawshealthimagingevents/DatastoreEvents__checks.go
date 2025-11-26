//go:build !no_runtime_type_checking

package previewawshealthimagingevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawshealthimaging"
)

func (d *jsiiProxy_DatastoreEvents) validateDataStoreCreatedPatternParameters(options *DatastoreEvents_DataStoreCreated_DataStoreCreatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreCreatingPatternParameters(options *DatastoreEvents_DataStoreCreating_DataStoreCreatingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreCreationFailedPatternParameters(options *DatastoreEvents_DataStoreCreationFailed_DataStoreCreationFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreDeletedPatternParameters(options *DatastoreEvents_DataStoreDeleted_DataStoreDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateDataStoreDeletingPatternParameters(options *DatastoreEvents_DataStoreDeleting_DataStoreDeletingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopiedPatternParameters(options *DatastoreEvents_ImageSetCopied_ImageSetCopiedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopyFailedPatternParameters(options *DatastoreEvents_ImageSetCopyFailed_ImageSetCopyFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopyingPatternParameters(options *DatastoreEvents_ImageSetCopying_ImageSetCopyingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCopyingWithReadOnlyAccessPatternParameters(options *DatastoreEvents_ImageSetCopyingWithReadOnlyAccess_ImageSetCopyingWithReadOnlyAccessProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetCreatedPatternParameters(options *DatastoreEvents_ImageSetCreated_ImageSetCreatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetDeletedPatternParameters(options *DatastoreEvents_ImageSetDeleted_ImageSetDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetDeletingPatternParameters(options *DatastoreEvents_ImageSetDeleting_ImageSetDeletingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetUpdatedPatternParameters(options *DatastoreEvents_ImageSetUpdated_ImageSetUpdatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetUpdateFailedPatternParameters(options *DatastoreEvents_ImageSetUpdateFailed_ImageSetUpdateFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImageSetUpdatingPatternParameters(options *DatastoreEvents_ImageSetUpdating_ImageSetUpdatingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobCompletedPatternParameters(options *DatastoreEvents_ImportJobCompleted_ImportJobCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobFailedPatternParameters(options *DatastoreEvents_ImportJobFailed_ImportJobFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobInProgressPatternParameters(options *DatastoreEvents_ImportJobInProgress_ImportJobInProgressProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DatastoreEvents) validateImportJobSubmittedPatternParameters(options *DatastoreEvents_ImportJobSubmitted_ImportJobSubmittedProps) error {
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

