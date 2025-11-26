package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.kms@AWSAPICallViaCloudTrail event types for Key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type KeyEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for KeyEvents_AWSAPICallViaCloudTrail
type jsiiProxy_KeyEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewKeyEvents_AWSAPICallViaCloudTrail() KeyEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_KeyEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKeyEvents_AWSAPICallViaCloudTrail_Override(k KeyEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		k,
	)
}

