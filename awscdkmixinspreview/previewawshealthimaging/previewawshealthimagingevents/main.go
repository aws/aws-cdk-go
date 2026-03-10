package previewawshealthimagingevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreated",
		reflect.TypeOf((*DataStoreCreated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DataStoreCreated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreated.DataStoreCreatedProps",
		reflect.TypeOf((*DataStoreCreated_DataStoreCreatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreating",
		reflect.TypeOf((*DataStoreCreating)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DataStoreCreating{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreating.DataStoreCreatingProps",
		reflect.TypeOf((*DataStoreCreating_DataStoreCreatingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreationFailed",
		reflect.TypeOf((*DataStoreCreationFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DataStoreCreationFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreCreationFailed.DataStoreCreationFailedProps",
		reflect.TypeOf((*DataStoreCreationFailed_DataStoreCreationFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleted",
		reflect.TypeOf((*DataStoreDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DataStoreDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleted.DataStoreDeletedProps",
		reflect.TypeOf((*DataStoreDeleted_DataStoreDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleting",
		reflect.TypeOf((*DataStoreDeleting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DataStoreDeleting{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DataStoreDeleting.DataStoreDeletingProps",
		reflect.TypeOf((*DataStoreDeleting_DataStoreDeletingProps)(nil)).Elem(),
	)
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
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopied",
		reflect.TypeOf((*ImageSetCopied)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetCopied{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopied.ImageSetCopiedProps",
		reflect.TypeOf((*ImageSetCopied_ImageSetCopiedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyFailed",
		reflect.TypeOf((*ImageSetCopyFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetCopyFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyFailed.ImageSetCopyFailedProps",
		reflect.TypeOf((*ImageSetCopyFailed_ImageSetCopyFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopying",
		reflect.TypeOf((*ImageSetCopying)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetCopying{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopying.ImageSetCopyingProps",
		reflect.TypeOf((*ImageSetCopying_ImageSetCopyingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyingWithReadOnlyAccess",
		reflect.TypeOf((*ImageSetCopyingWithReadOnlyAccess)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetCopyingWithReadOnlyAccess{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyingWithReadOnlyAccess.ImageSetCopyingWithReadOnlyAccessProps",
		reflect.TypeOf((*ImageSetCopyingWithReadOnlyAccess_ImageSetCopyingWithReadOnlyAccessProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCreated",
		reflect.TypeOf((*ImageSetCreated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetCreated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCreated.ImageSetCreatedProps",
		reflect.TypeOf((*ImageSetCreated_ImageSetCreatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleted",
		reflect.TypeOf((*ImageSetDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleted.ImageSetDeletedProps",
		reflect.TypeOf((*ImageSetDeleted_ImageSetDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleting",
		reflect.TypeOf((*ImageSetDeleting)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetDeleting{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleting.ImageSetDeletingProps",
		reflect.TypeOf((*ImageSetDeleting_ImageSetDeletingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdateFailed",
		reflect.TypeOf((*ImageSetUpdateFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetUpdateFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdateFailed.ImageSetUpdateFailedProps",
		reflect.TypeOf((*ImageSetUpdateFailed_ImageSetUpdateFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdated",
		reflect.TypeOf((*ImageSetUpdated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetUpdated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdated.ImageSetUpdatedProps",
		reflect.TypeOf((*ImageSetUpdated_ImageSetUpdatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdating",
		reflect.TypeOf((*ImageSetUpdating)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImageSetUpdating{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdating.ImageSetUpdatingProps",
		reflect.TypeOf((*ImageSetUpdating_ImageSetUpdatingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobCompleted",
		reflect.TypeOf((*ImportJobCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImportJobCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobCompleted.ImportJobCompletedProps",
		reflect.TypeOf((*ImportJobCompleted_ImportJobCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobFailed",
		reflect.TypeOf((*ImportJobFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImportJobFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobFailed.ImportJobFailedProps",
		reflect.TypeOf((*ImportJobFailed_ImportJobFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobInProgress",
		reflect.TypeOf((*ImportJobInProgress)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImportJobInProgress{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobInProgress.ImportJobInProgressProps",
		reflect.TypeOf((*ImportJobInProgress_ImportJobInProgressProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobSubmitted",
		reflect.TypeOf((*ImportJobSubmitted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ImportJobSubmitted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImportJobSubmitted.ImportJobSubmittedProps",
		reflect.TypeOf((*ImportJobSubmitted_ImportJobSubmittedProps)(nil)).Elem(),
	)
}
