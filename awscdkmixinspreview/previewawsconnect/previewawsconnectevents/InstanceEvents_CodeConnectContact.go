package previewawsconnectevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.connect@CodeConnectContact event types for Instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   codeConnectContact := #error#.NewCodeConnectContact()
//
// Experimental.
type InstanceEvents_CodeConnectContact interface {
}

// The jsii proxy struct for InstanceEvents_CodeConnectContact
type jsiiProxy_InstanceEvents_CodeConnectContact struct {
	_ byte // padding
}

// Experimental.
func NewInstanceEvents_CodeConnectContact() InstanceEvents_CodeConnectContact {
	_init_.Initialize()

	j := jsiiProxy_InstanceEvents_CodeConnectContact{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.CodeConnectContact",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceEvents_CodeConnectContact_Override(i InstanceEvents_CodeConnectContact) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.events.InstanceEvents.CodeConnectContact",
		nil, // no parameters
		i,
	)
}

