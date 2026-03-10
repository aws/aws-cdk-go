package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.kms@KMSImportedKeyMaterialExpiration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSImportedKeyMaterialExpiration := awscdkmixinspreview.Events.NewKMSImportedKeyMaterialExpiration()
//
// Experimental.
type KMSImportedKeyMaterialExpiration interface {
}

// The jsii proxy struct for KMSImportedKeyMaterialExpiration
type jsiiProxy_KMSImportedKeyMaterialExpiration struct {
	_ byte // padding
}

// Experimental.
func NewKMSImportedKeyMaterialExpiration() KMSImportedKeyMaterialExpiration {
	_init_.Initialize()

	j := jsiiProxy_KMSImportedKeyMaterialExpiration{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSImportedKeyMaterialExpiration",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKMSImportedKeyMaterialExpiration_Override(k KMSImportedKeyMaterialExpiration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSImportedKeyMaterialExpiration",
		nil, // no parameters
		k,
	)
}

// EventBridge event pattern for KMS Imported Key Material Expiration.
// Experimental.
func KMSImportedKeyMaterialExpiration_KmsImportedKeyMaterialExpirationPattern(options *KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateKMSImportedKeyMaterialExpiration_KmsImportedKeyMaterialExpirationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSImportedKeyMaterialExpiration",
		"kmsImportedKeyMaterialExpirationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

