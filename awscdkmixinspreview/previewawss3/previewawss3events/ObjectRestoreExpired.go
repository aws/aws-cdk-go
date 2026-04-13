package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectRestoreExpired.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRestoreExpired := awscdkmixinspreview.Events.NewObjectRestoreExpired()
//
// Experimental.
type ObjectRestoreExpired interface {
}

// The jsii proxy struct for ObjectRestoreExpired
type jsiiProxy_ObjectRestoreExpired struct {
	_ byte // padding
}

// Experimental.
func NewObjectRestoreExpired() ObjectRestoreExpired {
	_init_.Initialize()

	j := jsiiProxy_ObjectRestoreExpired{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectRestoreExpired_Override(o ObjectRestoreExpired) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Restore Expired.
// Experimental.
func ObjectRestoreExpired_EventPattern(options *ObjectRestoreExpired_ObjectRestoreExpiredProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectRestoreExpired_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

