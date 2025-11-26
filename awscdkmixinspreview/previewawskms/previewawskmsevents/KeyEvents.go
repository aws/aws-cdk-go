package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// EventBridge event patterns for Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyRef IKeyRef
//
//   keyEvents := awscdkmixinspreview.Events.KeyEvents_FromKey(keyRef)
//
// Experimental.
type KeyEvents interface {
	// EventBridge event pattern for Key AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *KeyEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Key KMS CMK Deletion.
	// Experimental.
	KMSCMKDeletionPattern(options *KeyEvents_KMSCMKDeletion_KMSCMKDeletionProps) *awsevents.EventPattern
	// EventBridge event pattern for Key KMS CMK Rotation.
	// Experimental.
	KMSCMKRotationPattern(options *KeyEvents_KMSCMKRotation_KMSCMKRotationProps) *awsevents.EventPattern
	// EventBridge event pattern for Key KMS Imported Key Material Expiration.
	// Experimental.
	KMSImportedKeyMaterialExpirationPattern(options *KeyEvents_KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps) *awsevents.EventPattern
}

// The jsii proxy struct for KeyEvents
type jsiiProxy_KeyEvents struct {
	_ byte // padding
}

// Create KeyEvents from a Key reference.
// Experimental.
func KeyEvents_FromKey(keyRef interfacesawskms.IKeyRef) KeyEvents {
	_init_.Initialize()

	if err := validateKeyEvents_FromKeyParameters(keyRef); err != nil {
		panic(err)
	}
	var returns KeyEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents",
		"fromKey",
		[]interface{}{keyRef},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyEvents) AwsAPICallViaCloudTrailPattern(options *KeyEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := k.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		k,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyEvents) KMSCMKDeletionPattern(options *KeyEvents_KMSCMKDeletion_KMSCMKDeletionProps) *awsevents.EventPattern {
	if err := k.validateKMSCMKDeletionPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		k,
		"kMSCMKDeletionPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyEvents) KMSCMKRotationPattern(options *KeyEvents_KMSCMKRotation_KMSCMKRotationProps) *awsevents.EventPattern {
	if err := k.validateKMSCMKRotationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		k,
		"kMSCMKRotationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyEvents) KMSImportedKeyMaterialExpirationPattern(options *KeyEvents_KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps) *awsevents.EventPattern {
	if err := k.validateKMSImportedKeyMaterialExpirationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		k,
		"kMSImportedKeyMaterialExpirationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

