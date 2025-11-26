package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.kms@KMSCMKRotation event types for Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSCMKRotation := #error#.NewKMSCMKRotation()
//
// Experimental.
type KeyEvents_KMSCMKRotation interface {
}

// The jsii proxy struct for KeyEvents_KMSCMKRotation
type jsiiProxy_KeyEvents_KMSCMKRotation struct {
	_ byte // padding
}

// Experimental.
func NewKeyEvents_KMSCMKRotation() KeyEvents_KMSCMKRotation {
	_init_.Initialize()

	j := jsiiProxy_KeyEvents_KMSCMKRotation{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKRotation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKeyEvents_KMSCMKRotation_Override(k KeyEvents_KMSCMKRotation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKRotation",
		nil, // no parameters
		k,
	)
}

