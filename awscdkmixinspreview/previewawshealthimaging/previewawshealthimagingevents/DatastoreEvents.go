package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawshealthimaging"
)

// EventBridge event patterns for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var datastoreRef IDatastoreRef
//
//   datastoreEvents := awscdkmixinspreview.Events.DatastoreEvents_FromDatastore(datastoreRef)
//
// Experimental.
type DatastoreEvents interface {
	// EventBridge event pattern for Datastore Data Store Created.
	// Experimental.
	DataStoreCreatedPattern(options *DatastoreEvents_DataStoreCreated_DataStoreCreatedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Data Store Creating.
	// Experimental.
	DataStoreCreatingPattern(options *DatastoreEvents_DataStoreCreating_DataStoreCreatingProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Data Store Creation Failed.
	// Experimental.
	DataStoreCreationFailedPattern(options *DatastoreEvents_DataStoreCreationFailed_DataStoreCreationFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Data Store Deleted.
	// Experimental.
	DataStoreDeletedPattern(options *DatastoreEvents_DataStoreDeleted_DataStoreDeletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Data Store Deleting.
	// Experimental.
	DataStoreDeletingPattern(options *DatastoreEvents_DataStoreDeleting_DataStoreDeletingProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Copied.
	// Experimental.
	ImageSetCopiedPattern(options *DatastoreEvents_ImageSetCopied_ImageSetCopiedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Copy Failed.
	// Experimental.
	ImageSetCopyFailedPattern(options *DatastoreEvents_ImageSetCopyFailed_ImageSetCopyFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Copying.
	// Experimental.
	ImageSetCopyingPattern(options *DatastoreEvents_ImageSetCopying_ImageSetCopyingProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Copying With Read Only Access.
	// Experimental.
	ImageSetCopyingWithReadOnlyAccessPattern(options *DatastoreEvents_ImageSetCopyingWithReadOnlyAccess_ImageSetCopyingWithReadOnlyAccessProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Created.
	// Experimental.
	ImageSetCreatedPattern(options *DatastoreEvents_ImageSetCreated_ImageSetCreatedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Deleted.
	// Experimental.
	ImageSetDeletedPattern(options *DatastoreEvents_ImageSetDeleted_ImageSetDeletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Deleting.
	// Experimental.
	ImageSetDeletingPattern(options *DatastoreEvents_ImageSetDeleting_ImageSetDeletingProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Updated.
	// Experimental.
	ImageSetUpdatedPattern(options *DatastoreEvents_ImageSetUpdated_ImageSetUpdatedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Update Failed.
	// Experimental.
	ImageSetUpdateFailedPattern(options *DatastoreEvents_ImageSetUpdateFailed_ImageSetUpdateFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Image Set Updating.
	// Experimental.
	ImageSetUpdatingPattern(options *DatastoreEvents_ImageSetUpdating_ImageSetUpdatingProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Import Job Completed.
	// Experimental.
	ImportJobCompletedPattern(options *DatastoreEvents_ImportJobCompleted_ImportJobCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Import Job Failed.
	// Experimental.
	ImportJobFailedPattern(options *DatastoreEvents_ImportJobFailed_ImportJobFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Import Job In Progress.
	// Experimental.
	ImportJobInProgressPattern(options *DatastoreEvents_ImportJobInProgress_ImportJobInProgressProps) *awsevents.EventPattern
	// EventBridge event pattern for Datastore Import Job Submitted.
	// Experimental.
	ImportJobSubmittedPattern(options *DatastoreEvents_ImportJobSubmitted_ImportJobSubmittedProps) *awsevents.EventPattern
}

// The jsii proxy struct for DatastoreEvents
type jsiiProxy_DatastoreEvents struct {
	_ byte // padding
}

// Create DatastoreEvents from a Datastore reference.
// Experimental.
func DatastoreEvents_FromDatastore(datastoreRef interfacesawshealthimaging.IDatastoreRef) DatastoreEvents {
	_init_.Initialize()

	if err := validateDatastoreEvents_FromDatastoreParameters(datastoreRef); err != nil {
		panic(err)
	}
	var returns DatastoreEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents",
		"fromDatastore",
		[]interface{}{datastoreRef},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) DataStoreCreatedPattern(options *DatastoreEvents_DataStoreCreated_DataStoreCreatedProps) *awsevents.EventPattern {
	if err := d.validateDataStoreCreatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"dataStoreCreatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) DataStoreCreatingPattern(options *DatastoreEvents_DataStoreCreating_DataStoreCreatingProps) *awsevents.EventPattern {
	if err := d.validateDataStoreCreatingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"dataStoreCreatingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) DataStoreCreationFailedPattern(options *DatastoreEvents_DataStoreCreationFailed_DataStoreCreationFailedProps) *awsevents.EventPattern {
	if err := d.validateDataStoreCreationFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"dataStoreCreationFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) DataStoreDeletedPattern(options *DatastoreEvents_DataStoreDeleted_DataStoreDeletedProps) *awsevents.EventPattern {
	if err := d.validateDataStoreDeletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"dataStoreDeletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) DataStoreDeletingPattern(options *DatastoreEvents_DataStoreDeleting_DataStoreDeletingProps) *awsevents.EventPattern {
	if err := d.validateDataStoreDeletingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"dataStoreDeletingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetCopiedPattern(options *DatastoreEvents_ImageSetCopied_ImageSetCopiedProps) *awsevents.EventPattern {
	if err := d.validateImageSetCopiedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetCopiedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetCopyFailedPattern(options *DatastoreEvents_ImageSetCopyFailed_ImageSetCopyFailedProps) *awsevents.EventPattern {
	if err := d.validateImageSetCopyFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetCopyFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetCopyingPattern(options *DatastoreEvents_ImageSetCopying_ImageSetCopyingProps) *awsevents.EventPattern {
	if err := d.validateImageSetCopyingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetCopyingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetCopyingWithReadOnlyAccessPattern(options *DatastoreEvents_ImageSetCopyingWithReadOnlyAccess_ImageSetCopyingWithReadOnlyAccessProps) *awsevents.EventPattern {
	if err := d.validateImageSetCopyingWithReadOnlyAccessPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetCopyingWithReadOnlyAccessPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetCreatedPattern(options *DatastoreEvents_ImageSetCreated_ImageSetCreatedProps) *awsevents.EventPattern {
	if err := d.validateImageSetCreatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetCreatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetDeletedPattern(options *DatastoreEvents_ImageSetDeleted_ImageSetDeletedProps) *awsevents.EventPattern {
	if err := d.validateImageSetDeletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetDeletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetDeletingPattern(options *DatastoreEvents_ImageSetDeleting_ImageSetDeletingProps) *awsevents.EventPattern {
	if err := d.validateImageSetDeletingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetDeletingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetUpdatedPattern(options *DatastoreEvents_ImageSetUpdated_ImageSetUpdatedProps) *awsevents.EventPattern {
	if err := d.validateImageSetUpdatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetUpdatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetUpdateFailedPattern(options *DatastoreEvents_ImageSetUpdateFailed_ImageSetUpdateFailedProps) *awsevents.EventPattern {
	if err := d.validateImageSetUpdateFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetUpdateFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImageSetUpdatingPattern(options *DatastoreEvents_ImageSetUpdating_ImageSetUpdatingProps) *awsevents.EventPattern {
	if err := d.validateImageSetUpdatingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"imageSetUpdatingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImportJobCompletedPattern(options *DatastoreEvents_ImportJobCompleted_ImportJobCompletedProps) *awsevents.EventPattern {
	if err := d.validateImportJobCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"importJobCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImportJobFailedPattern(options *DatastoreEvents_ImportJobFailed_ImportJobFailedProps) *awsevents.EventPattern {
	if err := d.validateImportJobFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"importJobFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImportJobInProgressPattern(options *DatastoreEvents_ImportJobInProgress_ImportJobInProgressProps) *awsevents.EventPattern {
	if err := d.validateImportJobInProgressPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"importJobInProgressPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastoreEvents) ImportJobSubmittedPattern(options *DatastoreEvents_ImportJobSubmitted_ImportJobSubmittedProps) *awsevents.EventPattern {
	if err := d.validateImportJobSubmittedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"importJobSubmittedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

