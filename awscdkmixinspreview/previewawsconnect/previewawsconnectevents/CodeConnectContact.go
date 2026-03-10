package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.connect@CodeConnectContact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeConnectContact := awscdkmixinspreview.Events.NewCodeConnectContact()
//
// Experimental.
type CodeConnectContact interface {
}

// The jsii proxy struct for CodeConnectContact
type jsiiProxy_CodeConnectContact struct {
	_ byte // padding
}

// Experimental.
func NewCodeConnectContact() CodeConnectContact {
	_init_.Initialize()

	j := jsiiProxy_CodeConnectContact{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCodeConnectContact_Override(c CodeConnectContact) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for Amazon Connect Contact Event.
// Experimental.
func CodeConnectContact_CodeConnectContactPattern(options *CodeConnectContact_CodeConnectContactProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCodeConnectContact_CodeConnectContactPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.events.CodeConnectContact",
		"codeConnectContactPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

