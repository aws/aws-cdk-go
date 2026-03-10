package previewawskmsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.kms@KMSCMKRotation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kMSCMKRotation := awscdkmixinspreview.Events.NewKMSCMKRotation()
//
// Experimental.
type KMSCMKRotation interface {
}

// The jsii proxy struct for KMSCMKRotation
type jsiiProxy_KMSCMKRotation struct {
	_ byte // padding
}

// Experimental.
func NewKMSCMKRotation() KMSCMKRotation {
	_init_.Initialize()

	j := jsiiProxy_KMSCMKRotation{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKRotation",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewKMSCMKRotation_Override(k KMSCMKRotation) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKRotation",
		nil, // no parameters
		k,
	)
}

// EventBridge event pattern for KMS CMK Rotation.
// Experimental.
func KMSCMKRotation_KmsCMKRotationPattern(options *KMSCMKRotation_KMSCMKRotationProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateKMSCMKRotation_KmsCMKRotationPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKRotation",
		"kmsCMKRotationPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

