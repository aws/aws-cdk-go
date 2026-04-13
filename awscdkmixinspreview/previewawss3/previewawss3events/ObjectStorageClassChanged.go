package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectStorageClassChanged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectStorageClassChanged := awscdkmixinspreview.Events.NewObjectStorageClassChanged()
//
// Experimental.
type ObjectStorageClassChanged interface {
}

// The jsii proxy struct for ObjectStorageClassChanged
type jsiiProxy_ObjectStorageClassChanged struct {
	_ byte // padding
}

// Experimental.
func NewObjectStorageClassChanged() ObjectStorageClassChanged {
	_init_.Initialize()

	j := jsiiProxy_ObjectStorageClassChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectStorageClassChanged_Override(o ObjectStorageClassChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Storage Class Changed.
// Experimental.
func ObjectStorageClassChanged_EventPattern(options *ObjectStorageClassChanged_ObjectStorageClassChangedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectStorageClassChanged_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

