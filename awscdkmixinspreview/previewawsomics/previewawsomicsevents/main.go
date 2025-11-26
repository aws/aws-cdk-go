package previewawsomicsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents",
		reflect.TypeOf((*ReferenceStoreEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "referenceImportJobStatusChangePattern", GoMethod: "ReferenceImportJobStatusChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "referenceStatusChangePattern", GoMethod: "ReferenceStatusChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_ReferenceStoreEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceImportJobStatusChange",
		reflect.TypeOf((*ReferenceStoreEvents_ReferenceImportJobStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ReferenceStoreEvents_ReferenceImportJobStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceImportJobStatusChange.ReferenceImportJobStatusChangeProps",
		reflect.TypeOf((*ReferenceStoreEvents_ReferenceImportJobStatusChange_ReferenceImportJobStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceStatusChange",
		reflect.TypeOf((*ReferenceStoreEvents_ReferenceStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ReferenceStoreEvents_ReferenceStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceStatusChange.ReferenceStatusChangeProps",
		reflect.TypeOf((*ReferenceStoreEvents_ReferenceStatusChange_ReferenceStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents",
		reflect.TypeOf((*SequenceStoreEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "readSetActivationJobStatusChangePattern", GoMethod: "ReadSetActivationJobStatusChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "readSetExportJobStatusChangePattern", GoMethod: "ReadSetExportJobStatusChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "readSetImportJobStatusChangePattern", GoMethod: "ReadSetImportJobStatusChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "readSetStatusChangePattern", GoMethod: "ReadSetStatusChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_SequenceStoreEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetActivationJobStatusChange",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetActivationJobStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SequenceStoreEvents_ReadSetActivationJobStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetActivationJobStatusChange.ReadSetActivationJobStatusChangeProps",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetActivationJobStatusChange_ReadSetActivationJobStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetExportJobStatusChange",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetExportJobStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SequenceStoreEvents_ReadSetExportJobStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetExportJobStatusChange.ReadSetExportJobStatusChangeProps",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetExportJobStatusChange_ReadSetExportJobStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetImportJobStatusChange",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetImportJobStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SequenceStoreEvents_ReadSetImportJobStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetImportJobStatusChange.ReadSetImportJobStatusChangeProps",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetImportJobStatusChange_ReadSetImportJobStatusChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetStatusChange",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetStatusChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_SequenceStoreEvents_ReadSetStatusChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetStatusChange.ReadSetStatusChangeProps",
		reflect.TypeOf((*SequenceStoreEvents_ReadSetStatusChange_ReadSetStatusChangeProps)(nil)).Elem(),
	)
}
