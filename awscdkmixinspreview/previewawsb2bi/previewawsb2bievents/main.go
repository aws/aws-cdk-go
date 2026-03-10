package previewawsb2bievents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted",
		reflect.TypeOf((*AcknowledgementCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AcknowledgementCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted.AckFileS3Attributes",
		reflect.TypeOf((*AcknowledgementCompleted_AckFileS3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted.AcknowledgementCompletedProps",
		reflect.TypeOf((*AcknowledgementCompleted_AcknowledgementCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted.InputFileS3Attributes",
		reflect.TypeOf((*AcknowledgementCompleted_InputFileS3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementFailed",
		reflect.TypeOf((*AcknowledgementFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AcknowledgementFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementFailed.AcknowledgementFailedProps",
		reflect.TypeOf((*AcknowledgementFailed_AcknowledgementFailedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementFailed.InputFileS3Attributes",
		reflect.TypeOf((*AcknowledgementFailed_InputFileS3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted",
		reflect.TypeOf((*TransformationCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TransformationCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted.InputFileS3Attributes",
		reflect.TypeOf((*TransformationCompleted_InputFileS3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted.OutputFileS3Attributes",
		reflect.TypeOf((*TransformationCompleted_OutputFileS3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted.TransformationCompletedProps",
		reflect.TypeOf((*TransformationCompleted_TransformationCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationFailed",
		reflect.TypeOf((*TransformationFailed)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TransformationFailed{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationFailed.InputFileS3Attributes",
		reflect.TypeOf((*TransformationFailed_InputFileS3Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationFailed.TransformationFailedProps",
		reflect.TypeOf((*TransformationFailed_TransformationFailedProps)(nil)).Elem(),
	)
}
