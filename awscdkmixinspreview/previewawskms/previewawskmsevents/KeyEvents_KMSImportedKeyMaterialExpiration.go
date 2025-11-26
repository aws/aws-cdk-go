package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.kms@KMSImportedKeyMaterialExpiration event types for Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSImportedKeyMaterialExpiration := #error#.NewKMSImportedKeyMaterialExpiration()
//
// Experimental.
type KeyEvents_KMSImportedKeyMaterialExpiration interface {
}

// The jsii proxy struct for KeyEvents_KMSImportedKeyMaterialExpiration
type jsiiProxy_KeyEvents_KMSImportedKeyMaterialExpiration struct {
	_ byte // padding
}

// Experimental.
func NewKeyEvents_KMSImportedKeyMaterialExpiration() KeyEvents_KMSImportedKeyMaterialExpiration {
	_init_.Initialize()

	j := jsiiProxy_KeyEvents_KMSImportedKeyMaterialExpiration{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSImportedKeyMaterialExpiration",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKeyEvents_KMSImportedKeyMaterialExpiration_Override(k KeyEvents_KMSImportedKeyMaterialExpiration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSImportedKeyMaterialExpiration",
		nil, // no parameters
		k,
	)
}

