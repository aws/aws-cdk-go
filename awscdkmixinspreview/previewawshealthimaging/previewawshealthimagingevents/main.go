package previewawshealthimagingevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents",
		reflect.TypeOf((*DatastoreEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "dataStoreCreatedPattern", GoMethod: "DataStoreCreatedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "dataStoreCreatingPattern", GoMethod: "DataStoreCreatingPattern"},
			_jsii_.MemberMethod{JsiiMethod: "dataStoreCreationFailedPattern", GoMethod: "DataStoreCreationFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "dataStoreDeletedPattern", GoMethod: "DataStoreDeletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "dataStoreDeletingPattern", GoMethod: "DataStoreDeletingPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetCopiedPattern", GoMethod: "ImageSetCopiedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetCopyFailedPattern", GoMethod: "ImageSetCopyFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetCopyingPattern", GoMethod: "ImageSetCopyingPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetCopyingWithReadOnlyAccessPattern", GoMethod: "ImageSetCopyingWithReadOnlyAccessPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetCreatedPattern", GoMethod: "ImageSetCreatedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetDeletedPattern", GoMethod: "ImageSetDeletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetDeletingPattern", GoMethod: "ImageSetDeletingPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetUpdatedPattern", GoMethod: "ImageSetUpdatedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetUpdateFailedPattern", GoMethod: "ImageSetUpdateFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "imageSetUpdatingPattern", GoMethod: "ImageSetUpdatingPattern"},
			_jsii_.MemberMethod{JsiiMethod: "importJobCompletedPattern", GoMethod: "ImportJobCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "importJobFailedPattern", GoMethod: "ImportJobFailedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "importJobInProgressPattern", GoMethod: "ImportJobInProgressPattern"},
			_jsii_.MemberMethod{JsiiMethod: "importJobSubmittedPattern", GoMethod: "ImportJobSubmittedPattern"},
		},
		func() interface{} {
			return &jsiiProxy_DatastoreEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreated",
		reflect.TypeOf((*DatastoreEvents_DataStoreCreated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_DataStoreCreated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreated.DataStoreCreatedProps",
		reflect.TypeOf((*DatastoreEvents_DataStoreCreated_DataStoreCreatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreating",
		reflect.TypeOf((*DatastoreEvents_DataStoreCreating)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_DataStoreCreating{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreating.DataStoreCreatingProps",
		reflect.TypeOf((*DatastoreEvents_DataStoreCreating_DataStoreCreatingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreationFailed",
		reflect.TypeOf((*DatastoreEvents_DataStoreCreationFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_DataStoreCreationFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreationFailed.DataStoreCreationFailedProps",
		reflect.TypeOf((*DatastoreEvents_DataStoreCreationFailed_DataStoreCreationFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleted",
		reflect.TypeOf((*DatastoreEvents_DataStoreDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_DataStoreDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleted.DataStoreDeletedProps",
		reflect.TypeOf((*DatastoreEvents_DataStoreDeleted_DataStoreDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleting",
		reflect.TypeOf((*DatastoreEvents_DataStoreDeleting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_DataStoreDeleting{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleting.DataStoreDeletingProps",
		reflect.TypeOf((*DatastoreEvents_DataStoreDeleting_DataStoreDeletingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopied",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopied)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetCopied{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopied.ImageSetCopiedProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopied_ImageSetCopiedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyFailed",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopyFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetCopyFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyFailed.ImageSetCopyFailedProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopyFailed_ImageSetCopyFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopying",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopying)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetCopying{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopying.ImageSetCopyingProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopying_ImageSetCopyingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyingWithReadOnlyAccess",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopyingWithReadOnlyAccess)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetCopyingWithReadOnlyAccess{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyingWithReadOnlyAccess.ImageSetCopyingWithReadOnlyAccessProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetCopyingWithReadOnlyAccess_ImageSetCopyingWithReadOnlyAccessProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCreated",
		reflect.TypeOf((*DatastoreEvents_ImageSetCreated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetCreated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCreated.ImageSetCreatedProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetCreated_ImageSetCreatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleted",
		reflect.TypeOf((*DatastoreEvents_ImageSetDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleted.ImageSetDeletedProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetDeleted_ImageSetDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleting",
		reflect.TypeOf((*DatastoreEvents_ImageSetDeleting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetDeleting{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleting.ImageSetDeletingProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetDeleting_ImageSetDeletingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdateFailed",
		reflect.TypeOf((*DatastoreEvents_ImageSetUpdateFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetUpdateFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdateFailed.ImageSetUpdateFailedProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetUpdateFailed_ImageSetUpdateFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdated",
		reflect.TypeOf((*DatastoreEvents_ImageSetUpdated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetUpdated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdated.ImageSetUpdatedProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetUpdated_ImageSetUpdatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdating",
		reflect.TypeOf((*DatastoreEvents_ImageSetUpdating)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImageSetUpdating{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdating.ImageSetUpdatingProps",
		reflect.TypeOf((*DatastoreEvents_ImageSetUpdating_ImageSetUpdatingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobCompleted",
		reflect.TypeOf((*DatastoreEvents_ImportJobCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImportJobCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobCompleted.ImportJobCompletedProps",
		reflect.TypeOf((*DatastoreEvents_ImportJobCompleted_ImportJobCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobFailed",
		reflect.TypeOf((*DatastoreEvents_ImportJobFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImportJobFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobFailed.ImportJobFailedProps",
		reflect.TypeOf((*DatastoreEvents_ImportJobFailed_ImportJobFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobInProgress",
		reflect.TypeOf((*DatastoreEvents_ImportJobInProgress)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImportJobInProgress{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobInProgress.ImportJobInProgressProps",
		reflect.TypeOf((*DatastoreEvents_ImportJobInProgress_ImportJobInProgressProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobSubmitted",
		reflect.TypeOf((*DatastoreEvents_ImportJobSubmitted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatastoreEvents_ImportJobSubmitted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobSubmitted.ImportJobSubmittedProps",
		reflect.TypeOf((*DatastoreEvents_ImportJobSubmitted_ImportJobSubmittedProps)(nil)).Elem(),
	)
}
