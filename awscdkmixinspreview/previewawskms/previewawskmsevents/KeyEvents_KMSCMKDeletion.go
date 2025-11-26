package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.kms@KMSCMKDeletion event types for Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSCMKDeletion := #error#.NewKMSCMKDeletion()
//
// Experimental.
type KeyEvents_KMSCMKDeletion interface {
}

// The jsii proxy struct for KeyEvents_KMSCMKDeletion
type jsiiProxy_KeyEvents_KMSCMKDeletion struct {
	_ byte // padding
}

// Experimental.
func NewKeyEvents_KMSCMKDeletion() KeyEvents_KMSCMKDeletion {
	_init_.Initialize()

	j := jsiiProxy_KeyEvents_KMSCMKDeletion{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKDeletion",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKeyEvents_KMSCMKDeletion_Override(k KeyEvents_KMSCMKDeletion) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKDeletion",
		nil, // no parameters
		k,
	)
}

